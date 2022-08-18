package threads

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/matg94/godirb/context"
	"github.com/matg94/godirb/data"
	"github.com/matg94/godirb/util"
)

func Worker(wg *sync.WaitGroup, appContext *context.AppContext, id int, client *http.Client) {
	for {
		request, err := appContext.RequestManager.GetNextRequest(appContext.Queue)
		if err == data.ErrEmptyQueue {
			appContext.DebugLogger.Low(id, "empty queue: ending thread")
			wg.Done()
			return
		}
		appContext.DebugLogger.Low(id, fmt.Sprintf("sending: %s/%s", request.URL, request.Path))
		code, err := request.Send(client)
		appContext.DebugLogger.Low(id, fmt.Sprintf("received: %d - %s/%s", code, request.URL, request.Path))

		if err != nil || code == -1 {
			appContext.RequestLogger.Log(code, request.URL, request.Path, false)
			appContext.DebugLogger.High(id, fmt.Sprintf("failed to send request, %s", err))
		} else if code == 404 || util.ListContains(code, appContext.AppConfig.WorkerConfig.IgnoreCodes) {
			appContext.RequestLogger.Log(code, request.URL, request.Path, false)
		} else {
			appContext.RequestLogger.Log(code, request.URL, request.Path, true)
		}
	}
}

func Start(appContext *context.AppContext) {
	var wg sync.WaitGroup
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for i := 0; i < appContext.AppConfig.WorkerConfig.Threads; i++ {
		wg.Add(1)
		go Worker(&wg, appContext, i, client)
	}
	wg.Wait()
}
