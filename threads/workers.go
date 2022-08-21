package threads

import (
	"net/http"
	"sync"
	"time"

	"github.com/matg94/godirb/context"
	"github.com/matg94/godirb/data"
	"github.com/matg94/godirb/requests"
	"github.com/matg94/godirb/util"
)

func Worker(wg *sync.WaitGroup, appContext *context.AppContext, id int, client *http.Client) {
	for {
		request, err := requests.GetNextRequest(appContext.Queue, appContext.BaseURL, appContext.AppConfig.RequestConfig)
		if err == data.ErrEmptyQueue {
			wg.Done()
			return
		}

		appContext.Limiter.AwaitPermission()
		code, err := request.Send(client)
		appContext.Limiter.Hit()
		appContext.ResultMap.Add(code)

		if err != nil || code == -1 {
			appContext.ErrorLogger.Log(&RequestLog{
				Code:    code,
				BaseURL: request.URL,
				Path:    request.Path,
			})
			appContext.DebugLogger.Log(&DebugLog{
				ThreadId: id,
				Location: "getting request response",
				Error:    err,
			})
		} else if code == 404 || util.ListContains(code, appContext.AppConfig.WorkerConfig.IgnoreCodes) {
			appContext.ErrorLogger.Log(&RequestLog{
				Code:    code,
				BaseURL: request.URL,
				Path:    request.Path,
			})
		} else {
			appContext.SuccessLogger.Log(&RequestLog{
				Code:    code,
				BaseURL: request.URL,
				Path:    request.Path,
			})
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
