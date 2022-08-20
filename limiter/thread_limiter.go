package limiter

import (
	"time"

	"github.com/matg94/godirb/timer"
)

type ThreadLimiter interface {
	AwaitPermission()
	CalculateCurrentRate() float64
	Hit()
}

type RequestThreadLimiter struct {
	TotalHits int
	Timer     timer.ProgramTimer
	MaxRate   float64
}

func CreateRequestLimiter(maxRate float64) *RequestThreadLimiter {
	timer := timer.CreateTimer()
	timer.Start()
	return &RequestThreadLimiter{
		TotalHits: 0,
		Timer:     timer,
		MaxRate:   maxRate,
	}
}

func (limiter *RequestThreadLimiter) AwaitPermission() {
	for limiter.CalculateCurrentRate() >= limiter.MaxRate {
		time.Sleep(time.Duration(1000/limiter.MaxRate) * time.Millisecond)
	}
}

func (limiter *RequestThreadLimiter) Hit() {
	limiter.TotalHits += 1
}

func (limiter *RequestThreadLimiter) CalculateCurrentRate() float64 {
	currentTime := limiter.Timer.GetCurrentTime()
	currentRate := float64(limiter.TotalHits) / currentTime.Seconds()
	return currentRate
}
