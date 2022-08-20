package context

import (
	"github.com/matg94/godirb/config"
	"github.com/matg94/godirb/data"
	"github.com/matg94/godirb/logger"
	"github.com/matg94/godirb/requests"
)

type AppContext struct {
	AppConfig      *config.AppConfig
	Queue          *data.WordQueue
	RequestManager *requests.RequestGenerator
	SuccessLogger  *logger.ThreadSafeLogger
	ErrorLogger    *logger.ThreadSafeLogger
	DebugLogger    *logger.ThreadSafeLogger
}
