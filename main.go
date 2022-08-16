package main

import (
	"github.com/matg94/godirb/config"
	"github.com/matg94/godirb/context"
	"github.com/matg94/godirb/data"
	"github.com/matg94/godirb/logger"
	"github.com/matg94/godirb/requests"
	"github.com/matg94/godirb/threads"
)

func main() {
	appConfig := config.LoadConfig("test")
	requestLogger := logger.CreateLogger(false, "")
	debugLogger := logger.CreateLogger(false, "")
	requestGenerator := &requests.RequestGenerator{
		BaseURL: "http://google.com",
	}
	wordQueue := data.CreateWordQueue()

	appContext := &context.AppContext{
		AppConfig:      appConfig,
		RequestManager: requestGenerator,
		Queue:          wordQueue,
		RequestLogger:  requestLogger,
		DebugLogger:    debugLogger,
	}

	config.ReadAndCompileWordLists(appContext.Queue, appConfig.WordLists, []string{})

	threads.Start(appContext)

	requestLogger.Output()
	debugLogger.Output()
}
