package threads

import (
	"fmt"
	"sync"

	"github.com/matg94/godirb/context"
	"github.com/matg94/godirb/data"
)

func Worker(wg *sync.WaitGroup, appContext *context.AppContext) {
	for {
		request, err := appContext.RequestManager.GetNextRequest(appContext.Queue)
		if err == data.ErrEmptyQueue {
			appContext.DebugLogger.Low("empty queue: ending thread")
			wg.Done()
			return
		}
		appContext.DebugLogger.High(fmt.Sprintf("sending: %s/%s", request.URL, request.Path))
		code, err := request.Send()
		if err != nil {
			appContext.DebugLogger.High(fmt.Sprintf("failed to send request, %s", err))
		}
		if code == 200 {
			appContext.RequestLogger.High(fmt.Sprintf("%d: %s/%s", code, request.URL, request.Path))
		} else {
			appContext.RequestLogger.Low(fmt.Sprintf("%d: %s/%s", code, request.URL, request.Path))
		}
	}
}

func Start(appContext *context.AppContext) {
	var wg sync.WaitGroup
	wg.Add(appContext.AppConfig.WorkerConfig.Threads)

	for i := 0; i < appContext.AppConfig.WorkerConfig.Threads; i++ {
		go Worker(&wg, appContext)
	}
	wg.Wait()
}
