package logger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/matg94/godirb/config"
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
	conf := config.LoggerTypeConfig{
		File:     "",
		Live:     false,
		JsonDump: false,
	}
	logger := CreateThreadSafeLogger(conf)
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

func TestOutputToFileJSON(t *testing.T) {
	conf := config.LoggerTypeConfig{
		File:     "testlog.json",
		Live:     false,
		JsonDump: false,
	}
	logger := CreateThreadSafeLogger(conf)
	testLog := &TestLog{
		Message: "test",
		Value:   1,
	}
	logger.Log(testLog)
	err := logger.Output()

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

// func TestOutputToFileInvalidFileGivesError(t *testing.T) {
// 	conf := config.LoggerTypeConfig{
// 		Enabled: true,
// 		File:    "",
// 		Json:    false,
// 		Live:    false,
// 	}
// 	logger := CreateThreadSafeLogger(conf)
// 	testLog := &TestLog{
// 		Message: "test",
// 		Value:   1,
// 	}
// 	logger.Log(testLog)
// 	err := logger.Output()

// 	if err != ErrInvalidFile {
// 		t.Log("expected invalid file error but got", err)
// 		t.Fail()
// 	}
// }
