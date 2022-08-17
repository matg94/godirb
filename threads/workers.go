package threads

import (
	"fmt"
	"sync"

	"github.com/matg94/godirb/context"
	"github.com/matg94/godirb/data"
)

func Worker(wg *sync.WaitGroup, appContext *context.AppContext, id int) {
	for {
		request, err := appContext.RequestManager.GetNextRequest(appContext.Queue)
		if err == data.ErrEmptyQueue {
			appContext.DebugLogger.Low(id, "empty queue: ending thread")
			wg.Done()
			return
		}
		appContext.DebugLogger.Low(id, fmt.Sprintf("sending: %s/%s", request.URL, request.Path))
		code, err := request.Send()
		appContext.DebugLogger.Low(id, fmt.Sprintf("received: %d - %s/%s", code, request.URL, request.Path))
		if err != nil || code == -1 {
			appContext.RequestLogger.Log(code, request.URL, request.Path)
			appContext.DebugLogger.High(id, fmt.Sprintf("failed to send request, %s", err))
		}
		// TODO: move this logic to the output.go
		appContext.RequestLogger.Log(code, request.URL, request.Path)
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
