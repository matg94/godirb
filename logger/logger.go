package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Log struct {
	content   string
	timestamp string
	debug     bool
}

func (log *Log) toString() string {
	return fmt.Sprintf("%s: %s", log.timestamp, log.content)
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

func (log *Logger) Low(message string) {
	log.Logs = append(log.Logs, Log{
		content:   message,
		timestamp: time.Now().Local().Format(time.RFC822Z),
		debug:     true,
	})
}

func (log *Logger) High(message string) {
	log.Logs = append(log.Logs, Log{
		content:   message,
		timestamp: time.Now().Local().Format(time.RFC822Z),
		debug:     false,
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
			if !log.Debug && lg.debug {
				continue
			}
			output += lg.toString() + "\n"
		}
		ioutil.WriteFile(log.OutputFile, []byte(output), 0644)
	} else {
		for _, lg := range log.Logs {
			if !log.Debug && lg.debug {
				continue
			}
			fmt.Println(lg.toString())
		}
	}
	return nil
}
