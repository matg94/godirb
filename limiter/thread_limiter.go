package limiter

import (
	"time"

	"github.com/matg94/godirb/timer"
)

type ThreadLimiter interface {
	AwaitPermission() bool
	CalculateCurrentRate() float64
	Hit()
}

type RequestThreadLimiter struct {
	Enabled   bool
	TotalHits int
	Timer     timer.ProgramTimer
	MaxRate   float64
}

func CreateRequestLimiter(maxRate float64, enabled bool) *RequestThreadLimiter {
	timer := timer.CreateTimer()
	timer.Start()
	return &RequestThreadLimiter{
		Enabled:   enabled,
		TotalHits: 0,
		Timer:     timer,
		MaxRate:   maxRate,
	}
}

func (limiter *RequestThreadLimiter) AwaitPermission() bool {
	if !limiter.Enabled {
		return false
	}
	for limiter.CalculateCurrentRate() >= limiter.MaxRate {
		time.Sleep(time.Duration(1000/limiter.MaxRate) * time.Millisecond)
	}
	return true
}

func (limiter *RequestThreadLimiter) Hit() {
	limiter.TotalHits += 1
}

func (limiter *RequestThreadLimiter) CalculateCurrentRate() float64 {
	currentTime := limiter.Timer.GetCurrentTime()
	currentRate := float64(limiter.TotalHits) / currentTime.Seconds()
	return currentRate
}
