package timer

import "time"

type ProgramTimer interface {
	Start()
	Stop()
	GetRunTime() time.Duration
	GetCurrentTime() time.Duration
}

type Timer struct {
	startTime time.Time
	endTime   time.Time
}

func CreateTimer() *Timer {
	return &Timer{}
}

func (t *Timer) Start() {
	t.startTime = time.Now()
}

func (t *Timer) Stop() {
	t.endTime = time.Now()
}

func (t *Timer) GetRunTime() time.Duration {
	return t.endTime.Sub(t.startTime)
}

func (t *Timer) GetCurrentTime() time.Duration {
	return time.Now().Sub(t.startTime)
}
