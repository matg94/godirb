package context

import (
	"github.com/matg94/godirb/config"
	"github.com/matg94/godirb/data"
	"github.com/matg94/godirb/limiter"
	"github.com/matg94/godirb/logger"
)

type AppContext struct {
	AppConfig     *config.AppConfig
	BaseURL       string
	Queue         *data.WordQueue
	Limiter       *limiter.RequestThreadLimiter
	SuccessLogger *logger.ThreadSafeLogger
	ErrorLogger   *logger.ThreadSafeLogger
	DebugLogger   *logger.ThreadSafeLogger
	ResultMap     *logger.RequestCounterMap
}
