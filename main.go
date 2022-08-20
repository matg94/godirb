package main

import (
	"fmt"

	"github.com/matg94/godirb/config"
	"github.com/matg94/godirb/context"
	"github.com/matg94/godirb/data"
	"github.com/matg94/godirb/logger"
	"github.com/matg94/godirb/requests"
	"github.com/matg94/godirb/threads"
	"github.com/matg94/godirb/timer"
)

func CreateLoggers(config *config.AppConfig) (*logger.ThreadSafeLogger, *logger.ThreadSafeLogger, *logger.ThreadSafeLogger) {
	successLogger := logger.CreateThreadSafeLogger(config.LoggingConfig.SuccessLogger)
	errorLogger := logger.CreateThreadSafeLogger(config.LoggingConfig.ErrorLogger)
	debugLogger := logger.CreateThreadSafeLogger(config.LoggingConfig.DebugLogger)
	return successLogger, errorLogger, debugLogger
}

func main() {
	appConfig := config.LoadConfig("test")
	successLogger, errorLogger, debugLogger := CreateLoggers(appConfig)
	requestGenerator := &requests.RequestGenerator{
		BaseURL: "http://localhost",
	}
	wordQueue := data.CreateWordQueue()

	appContext := &context.AppContext{
		AppConfig:      appConfig,
		RequestManager: requestGenerator,
		Queue:          wordQueue,
		SuccessLogger:  successLogger,
		ErrorLogger:    errorLogger,
		DebugLogger:    debugLogger,
	}

	config.ReadAndCompileWordLists(appContext.Queue, appConfig.WorkerConfig.WordLists, []string{}, appContext.AppConfig.WorkerConfig.Append)

	if appConfig.LoggingConfig.Stats {
		fmt.Println("-------------------------------")
		fmt.Println("Words Generated: ", len(wordQueue.GetAll()))
		fmt.Println("-------------------------------")
	}

	mainTimer := timer.CreateTimer()
	mainTimer.Start()
	threads.Start(appContext)

	mainTimer.Stop()

	if appConfig.LoggingConfig.Stats {
		fmt.Println("-------------------------------")
		fmt.Println("Time taken:", mainTimer.GetRunTime().Seconds(), "seconds")
		fmt.Println("-------------------------------")
	}
	debugLogger.Output()
	successLogger.Output()
	errorLogger.Output()
}

// func (out *Outputter) GetStats() map[int]int {
// 	responses := map[int]int{}
// 	for _, res := range out.Results {
// 		responses[res.Code] += 1
// 	}
// 	return responses
// }
