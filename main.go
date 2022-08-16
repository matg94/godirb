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
	requestLogger := logger.CreateLogger(appConfig.LoggingConfig.Debug, "")
	debugLogger := logger.CreateLogger(appConfig.LoggingConfig.Debug, "log.txt")
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

	config.ReadAndCompileWordLists(appContext.Queue, appConfig.WorkerConfig.WordLists, []string{}, []string{})

	threads.Start(appContext)

	requestLogger.Output()
	debugLogger.Output()
}
