package limiter

import (
	"testing"
	"time"
)

type MockProgramTimer struct {
	GetTimeReturn time.Duration
}

func (t *MockProgramTimer) Start() {}

func (t *MockProgramTimer) Stop() {}

func (t *MockProgramTimer) GetRunTime() time.Duration {
	return time.Duration(1)
}

func (t *MockProgramTimer) GetCurrentTime() time.Duration {
	return t.GetTimeReturn
}

func TestAwaitingPermission(t *testing.T) {
	lmter := CreateRequestLimiter(2, true)
	lmter.Timer = &MockProgramTimer{
		GetTimeReturn: time.Duration(3 * time.Second),
	}
	lmter.TotalHits = 1
	threadSlept := lmter.AwaitPermission()
	if !threadSlept {
		t.Log("expected thread slept true but got", threadSlept)
		t.Fail()
	}
}

func TestAwaitingPermissionDisabled(t *testing.T) {
	lmter := CreateRequestLimiter(2, false)
	lmter.TotalHits = 1
	threadSlept := lmter.AwaitPermission()
	if threadSlept {
		t.Log("expected thread slept false but got", threadSlept)
		t.Fail()
	}
}
