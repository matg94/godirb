package config

import (
	"reflect"
	"testing"

	"github.com/matg94/godirb/data"
)

func TestReadAndCompileWordlistsWithAppend(t *testing.T) {
	queue := data.CreateWordQueue()
	ReadAndCompileWordLists(
		queue,
		[]string{"./test.txt"},
		[]string{"hello2"},
		[]string{".html"},
	)

	expected := []string{
		"hello2.html",
		"hello.html",
		"echo.html",
		"test.html",
	}

	result := queue.GetAll()

	if !reflect.DeepEqual(result, expected) {
		t.Log("expected queue to contain ", expected, "but got", result)
		t.Fail()
	}
}

func TestReadAndCompileWordlistsWithoutAppend(t *testing.T) {
	queue := data.CreateWordQueue()
	ReadAndCompileWordLists(
		queue,
		[]string{"./test.txt"},
		[]string{"hello2"},
		[]string{},
	)

	expected := []string{
		"hello2",
		"hello",
		"echo",
		"test",
	}

	result := queue.GetAll()

	if !reflect.DeepEqual(result, expected) {
		t.Log("expected queue to contain ", expected, "but got", result)
		t.Fail()
	}
}
