package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type Loggable interface {
	ToString() string
	ToJSON() string
}

var ErrInvalidFile error = errors.New("could not write to given file")

type ThreadSafeLogger struct {
	Logs       []Loggable
	OutputLive bool
	mut        sync.Mutex
}

func CreateThreadSafeLogger(live bool) *ThreadSafeLogger {
	return &ThreadSafeLogger{
		Logs:       []Loggable{},
		OutputLive: live,
		mut:        sync.Mutex{},
	}
}

func (logger *ThreadSafeLogger) Output(asJson bool) {
	out := os.Stdout
	for _, log := range logger.Logs {
		if asJson {
			fmt.Fprintln(out, log.ToJSON())
			continue
		}
		fmt.Fprintln(out, log.ToString())
	}
}

func (logger *ThreadSafeLogger) OutputToFile(asJson bool, filepath string) error { // TODO: Add JSON functionality (make whole file json)
	if filepath != "" {
		f, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer f.Close()

		if asJson {
			jsonString, err := json.Marshal(logger.Logs)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(filepath, jsonString, 0644)
			if err != nil {
				return err
			}
		} else {
			var output string
			for _, lg := range logger.Logs {
				output += lg.ToString() + "\n"
			}
			err := ioutil.WriteFile(filepath, []byte(output), 0644)
			if err != nil {
				return err
			}
		}
	} else {
		return ErrInvalidFile
	}
	return nil
}

func (logger *ThreadSafeLogger) Log(log Loggable) {
	logger.mut.Lock()
	defer logger.mut.Unlock()

	logger.Logs = append(logger.Logs, log)
	if !logger.OutputLive {
		return
	}
	fmt.Println(log.ToString())
}
