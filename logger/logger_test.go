package logger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type TestLog struct {
	Message string `json:"message"`
	Value   int    `json:"value"`
}

func (t *TestLog) ToJSON() string {
	jsonString, _ := json.Marshal(t)
	return string(jsonString)
}

func (t *TestLog) ToString() string {
	return fmt.Sprintf("%d | %s", t.Value, t.Message)
}

func ReadFile(filepath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func TestAddLog(t *testing.T) {
	logger := CreateThreadSafeLogger(false)
	testLog := &TestLog{
		Message: "test",
		Value:   1,
	}
	logger.Log(testLog)
	if logger.Logs[0].ToString() != testLog.ToString() {
		t.Logf("expected first log to be equal to %s, but got %s", testLog.ToString(), logger.Logs[0].ToString())
		t.Fail()
	}
}

func TestOutputToFile(t *testing.T) {
	logger := CreateThreadSafeLogger(false)
	testLog := &TestLog{
		Message: "test",
		Value:   1,
	}
	logger.Log(testLog)
	err := logger.OutputToFile(false, "testlog.txt")

	if err != nil {
		t.Log("expected no errors but got", err)
		t.Fail()
	}

	data, err := ReadFile("testlog.txt")
	if err != nil {
		t.Log("expected no errors reading file but got", err)
		t.Fail()
	}

	if string(data) != (testLog.ToString() + "\n") {
		t.Logf("expected file contents to match '%s' (plus new line) but got '%s'", testLog.ToString(), string(data))
		t.Fail()
	}
	os.Remove("testlog.txt")
}

func TestOutputToFileJSON(t *testing.T) {
	logger := CreateThreadSafeLogger(false)
	testLog := &TestLog{
		Message: "test",
		Value:   1,
	}
	logger.Log(testLog)
	err := logger.OutputToFile(true, "testlog.json")

	if err != nil {
		t.Log("expected no errors but got", err)
		t.Fail()
	}

	data, err := ReadFile("testlog.json")
	if err != nil {
		t.Log("expected no errors reading file but got", err)
		t.Fail()
	}

	if string(data) != "[{\"message\":\"test\",\"value\":1}]" {
		t.Logf("expected file contents to match '%s' (plus new line) but got '%s'", testLog.ToString(), string(data))
		t.Fail()
	}
	os.Remove("testlog.json")
}

func TestOutputToFileInvalidFileGivesError(t *testing.T) {
	logger := CreateThreadSafeLogger(false)
	testLog := &TestLog{
		Message: "test",
		Value:   1,
	}
	logger.Log(testLog)
	err := logger.OutputToFile(false, "")

	if err != ErrInvalidFile {
		t.Log("expected invalid file error but got", err)
		t.Fail()
	}
}
