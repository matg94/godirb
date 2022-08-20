package logger

import (
	"bytes"
	"testing"
)

func TestThreadMapAdd(t *testing.T) {
	threadMap := CreateRequestCounterMap()
	threadMap.Add(1)
	threadMap.Add(1)
	if threadMap.Map[1] != 2 {
		t.Log("expected threadmap [1] to contain 2 but got", threadMap.Map[1])
		t.Fail()
	}
}

func TestThreadMapPrint(t *testing.T) {
	threadMap := CreateRequestCounterMap()
	threadMap.Add(1)
	threadMap.Add(1)

	var b bytes.Buffer
	threadMap.Print(&b)
	got := b.String()
	expected := "Code  | Count\n----- | -----\n1     | 2\n"
	if got != expected {
		t.Logf("expected output to be %q, but got %q", expected, got)
		t.Fail()
	}
}
