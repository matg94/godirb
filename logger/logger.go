package logger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Log struct {
	ThreadId  int    `json:"threadId"`
	Content   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Debug     bool   `json:"levelDebug"`
}

func (log *Log) toJSON() string {
	res, err := json.Marshal(&log)
	if err != nil {
		return "{\"error\": \"parsing json logs error\"}"
	}
	return string(res)
}

type Logger struct {
	Debug      bool
	OutputFile string
	Logs       []Log
}

func CreateLogger(debug bool, outputFile string) *Logger {
	return &Logger{
		Debug:      debug,
		OutputFile: outputFile,
		Logs:       []Log{},
	}
}

func (log *Logger) Low(thread int, message string) {
	log.Logs = append(log.Logs, Log{
		ThreadId:  thread,
		Content:   message,
		Timestamp: time.Now().Local().Format(time.RFC822Z),
		Debug:     true,
	})
}

func (log *Logger) High(thread int, message string) {
	log.Logs = append(log.Logs, Log{
		ThreadId:  thread,
		Content:   message,
		Timestamp: time.Now().Local().Format(time.RFC822Z),
		Debug:     false,
	})
}

func (log *Logger) Output() error {
	if log.OutputFile != "" {
		file, err := os.Create(log.OutputFile)
		if err != nil {
			return err
		}
		defer file.Close()
		var output string
		for _, lg := range log.Logs {
			if !log.Debug && lg.Debug {
				continue
			}
			output += lg.toJSON() + "\n"
		}
		ioutil.WriteFile(log.OutputFile, []byte(output), 0644)
	} else {
		for _, lg := range log.Logs {
			if !log.Debug && lg.Debug {
				continue
			}
			fmt.Println(lg.toJSON())
		}
	}
	return nil
}
