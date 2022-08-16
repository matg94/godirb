package data

import "testing"

func TestQueueAdd(t *testing.T) {
	queue := CreateWordQueue()
	queue.AddList([]string{"1", "2"})
	next, err := queue.Next()
	if next != "1" || err != nil {
		t.Log("expected 1 as next value and no errors but got", next, err)
		t.Fail()
	}
	next, err = queue.Next()
	if next != "2" || err != nil {
		t.Log("expected 2 as next value and no errors but got", next, err)
		t.Fail()
	}
	next, err = queue.Next()
	if next != "" || err != ErrEmptyQueue {
		t.Log("expected empty string and queue err but got", next, err)
		t.Fail()
	}
}
