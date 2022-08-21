package requests

import (
	"testing"

	"github.com/matg94/godirb/config"
	"github.com/matg94/godirb/data"
)

func TestGetNextRequest(t *testing.T) {
	queue := data.CreateWordQueue()
	config.ReadAndCompileWordLists(
		queue,
		[]string{"../config/test.txt"},
		[]string{},
		[]string{},
		true,
	)

	req, err := GetNextRequest(queue, "localhost", config.RequestConfig{})
	if err != nil {
		t.Log("expected no errors but got", err)
		t.Fail()
	}

	if req.Path != "hello" {
		t.Log("expected first path to be 'hello' but got", req.Path)
		t.Fail()
	}

	req, err = GetNextRequest(queue, "localhost", config.RequestConfig{})
	if err != nil {
		t.Log("expected no errors but got", err)
		t.Fail()
	}

	if req.Path != "echo" {
		t.Log("expected second path to be 'echo' but got", req.Path)
		t.Fail()
	}

	if len(queue.GetAll()) != 1 {
		t.Log("expected queue to have one word left but got", len(queue.GetAll()))
		t.Fail()
	}
}

func TestGetNextRequestErrEmptyQueue(t *testing.T) {
	queue := data.CreateWordQueue()
	config.ReadAndCompileWordLists(
		queue,
		[]string{},
		[]string{},
		[]string{},
		true,
	)

	_, err := GetNextRequest(queue, "localhost", config.RequestConfig{})
	if err != data.ErrEmptyQueue {
		t.Log("expected err empty queue but got", err)
		t.Fail()
	}

}
