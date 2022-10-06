package main

import (
	"fmt"
	"log"
	"os"

	"github.com/matg94/godirb/config"
	"github.com/matg94/godirb/context"
	"github.com/matg94/godirb/data"
	"github.com/matg94/godirb/limiter"
	"github.com/matg94/godirb/logger"
	"github.com/matg94/godirb/threads"
	"github.com/matg94/godirb/timer"
)

var version string = "v1.0"

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
	parsedFlags := ParseFlags()
	if parsedFlags.Version {
		fmt.Println(version)
		os.Exit(0)
	}
	appConfig := config.LoadConfigWithFlags(parsedFlags)
	successLogger, errorLogger, debugLogger := CreateLoggers(appConfig)
	wordQueue := data.CreateWordQueue()
	requestLimiter := limiter.CreateRequestLimiter(appConfig.WorkerConfig.Limiter.RequestsPerSecond, appConfig.WorkerConfig.Limiter.Enabled)
	threadSafeMap := logger.CreateRequestCounterMap()

	if parsedFlags.URL == "" {
		log.Fatal("no url provided")
		os.Exit(0)
	}

	appContext := &context.AppContext{
		AppConfig:     appConfig,
		BaseURL:       parsedFlags.URL,
		Queue:         wordQueue,
		Limiter:       requestLimiter,
		SuccessLogger: successLogger,
		ErrorLogger:   errorLogger,
		DebugLogger:   debugLogger,
		ResultMap:     threadSafeMap,
	}

	config.ReadAndCompileWordLists(
		appContext.Queue,
		appConfig.WorkerConfig.WordLists,
		[]string{},
		appContext.AppConfig.WorkerConfig.Append,
		appConfig.WorkerConfig.AppendOnly,
		len(parsedFlags.Wordlist) > 0,
	)

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
		appContext.ResultMap.Print(os.Stdout)
		fmt.Println("-------------------------------")
		fmt.Println("Time taken:", mainTimer.GetRunTime().Seconds(), "seconds")
		fmt.Println("Total Hits:", int(requestLimiter.TotalHits))
		fmt.Println("Final Rate:", int(requestLimiter.CalculateCurrentRate()), "requests per second")
		fmt.Println("-------------------------------")
	}
	err := debugLogger.Output(os.Stdout)
	HandleFatalErr(err)
	err = successLogger.Output(os.Stdout)
	HandleFatalErr(err)
	err = errorLogger.Output(os.Stdout)
	HandleFatalErr(err)

}
