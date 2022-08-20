package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/matg94/godirb/config"
)

type Loggable interface {
	ToString() string
	ToJSON() string
}

var ErrInvalidFile error = errors.New("could not write to given file")

type ThreadSafeLogger struct {
	Logs     []Loggable
	JsonDump bool
	Live     bool
	Filepath string
	mut      sync.Mutex
}

func CreateThreadSafeLogger(config config.LoggerTypeConfig) *ThreadSafeLogger {
	return &ThreadSafeLogger{
		Logs:     []Loggable{},
		JsonDump: config.JsonDump,
		Live:     config.Live,
		Filepath: config.File,
		mut:      sync.Mutex{},
	}
}

func (logger *ThreadSafeLogger) OutputString(writer io.Writer) {
	for _, log := range logger.Logs {
		fmt.Fprintln(writer, log.ToString())
	}
}

func (logger *ThreadSafeLogger) OutputJSON(writer io.Writer) error {
	jsonString, err := json.Marshal(logger.Logs)
	if err != nil {
		return err
	}
	fmt.Fprint(writer, string(jsonString))
	return nil
}

func (logger *ThreadSafeLogger) Output() error {
	if logger.Filepath != "" {
		file, err := os.Create(logger.Filepath)
		if err != nil {
			return err
		}
		defer file.Close()

		err = logger.OutputJSON(file)
		if err != nil {
			return err
		}
	}
	if logger.JsonDump {
		logger.OutputJSON(os.Stdout)
	}

	return nil
}

func (logger *ThreadSafeLogger) Log(log Loggable) {
	logger.mut.Lock()
	defer logger.mut.Unlock()

	logger.Logs = append(logger.Logs, log)
	if !logger.Live {
		return
	}
	fmt.Println(log.ToString())
}
