package main

import (
	"fmt"

	"github.com/matg94/godirb/config"
	"github.com/matg94/godirb/context"
	"github.com/matg94/godirb/data"
	"github.com/matg94/godirb/logger"
	"github.com/matg94/godirb/requests"
	"github.com/matg94/godirb/threads"
)

func main() {
	appConfig := config.LoadConfig("test")
	requestLogger := logger.CreateOutput(appConfig.LoggingConfig.DisplayLive)
	debugLogger := logger.CreateLogger(appConfig.LoggingConfig.Debug, "log.txt")
	requestGenerator := &requests.RequestGenerator{
		BaseURL: "http://localhost",
	}
	wordQueue := data.CreateWordQueue()

	appContext := &context.AppContext{
		AppConfig:      appConfig,
		RequestManager: requestGenerator,
		Queue:          wordQueue,
		RequestLogger:  requestLogger,
		DebugLogger:    debugLogger,
	}

	config.ReadAndCompileWordLists(appContext.Queue, appConfig.WorkerConfig.WordLists, []string{}, appContext.AppConfig.WorkerConfig.Append)

	// fmt.Println(wordQueue.GetAll())

	fmt.Println("-------------------------------")
	fmt.Println("Words Generated: ", len(wordQueue.GetAll()))
	fmt.Println("-------------------------------")

	threads.Start(appContext)

	fmt.Println("-------------------------------")
	requestLogger.OutputPretty()
	fmt.Println("-------------------------------")
	debugLogger.Output()
}
