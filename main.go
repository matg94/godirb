package main

import (
	"fmt"
	"log"

	"github.com/matg94/godirb/config"
	"github.com/matg94/godirb/context"
	"github.com/matg94/godirb/data"
	"github.com/matg94/godirb/limiter"
	"github.com/matg94/godirb/logger"
	"github.com/matg94/godirb/requests"
	"github.com/matg94/godirb/threads"
	"github.com/matg94/godirb/timer"
)

func HandleFatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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
	requestLimiter := limiter.CreateRequestLimiter(appConfig.WorkerConfig.Limiter.RequestsPerSecond, appConfig.WorkerConfig.Limiter.Enabled)
	threadSafeMap := logger.CreateRequestCounterMap()

	appContext := &context.AppContext{
		AppConfig:      appConfig,
		RequestManager: requestGenerator,
		Queue:          wordQueue,
		Limiter:        requestLimiter,
		SuccessLogger:  successLogger,
		ErrorLogger:    errorLogger,
		DebugLogger:    debugLogger,
		ResultMap:      threadSafeMap,
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
		appContext.ResultMap.Print()
		fmt.Println("-------------------------------")
		fmt.Println("Time taken:", mainTimer.GetRunTime().Seconds(), "seconds")
		fmt.Println("Total Hits:", int(requestLimiter.TotalHits))
		fmt.Println("Final Rate:", int(requestLimiter.CalculateCurrentRate()), "requests per second")
		fmt.Println("-------------------------------")
	}
	err := debugLogger.Output()
	HandleFatalErr(err)
	err = successLogger.Output()
	HandleFatalErr(err)
	err = errorLogger.Output()
	HandleFatalErr(err)

}
