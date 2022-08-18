package threads

import (
	"fmt"
	"sync"

	"github.com/matg94/godirb/context"
	"github.com/matg94/godirb/data"
	"github.com/matg94/godirb/util"
)

func Worker(wg *sync.WaitGroup, appContext *context.AppContext, id int) {
	for {
		if id == 3000 {
			fmt.Println("Getting request")
		}
		request, err := appContext.RequestManager.GetNextRequest(appContext.Queue)
		if err == data.ErrEmptyQueue {
			if id == 3000 {
				fmt.Println("empty")
			}
			appContext.DebugLogger.Low(id, "empty queue: ending thread")
			wg.Done()
			return
		}
		appContext.DebugLogger.Low(id, fmt.Sprintf("sending: %s/%s", request.URL, request.Path))
		if id == 3000 {
			fmt.Println("sending request")
		}
		code, err := request.Send()
		if id == 3000 {
			fmt.Println("got ", code)
		}
		appContext.DebugLogger.Low(id, fmt.Sprintf("received: %d - %s/%s", code, request.URL, request.Path))
		if err != nil || code == -1 {
			appContext.RequestLogger.Log(code, request.URL, request.Path, false)
			appContext.DebugLogger.High(id, fmt.Sprintf("failed to send request, %s", err))
		}
		if code == 404 || util.ListContains(code, appContext.AppConfig.WorkerConfig.IgnoreCodes) {
			appContext.RequestLogger.Log(code, request.URL, request.Path, false)
		} else {
			appContext.RequestLogger.Log(code, request.URL, request.Path, true)
		}
	}
}

func Start(appContext *context.AppContext) {
	var wg sync.WaitGroup
	wg.Add(appContext.AppConfig.WorkerConfig.Threads)

	for i := 0; i < appContext.AppConfig.WorkerConfig.Threads; i++ {
		go Worker(&wg, appContext, i+1)
	}
	wg.Wait()
}
