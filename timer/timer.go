package timer

import "time"

type ProgramTimer interface {
	Start()
	Stop()
	GetRunTime() time.Duration
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
