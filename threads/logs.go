package threads

import (
	"encoding/json"
	"fmt"
)

type DebugLog struct {
	Error    error  `json:"error"`
	ThreadId int    `json:"threadId"`
	Location string `json:"location"`
}

func (log *DebugLog) ToString() string {
	return fmt.Sprintf("ERROR - Thread: %d | Location: %s | Error: %s", log.ThreadId, log.Location, log.Error.Error())
}

func (log *DebugLog) ToJSON() string {
	jsonString, _ := json.Marshal(log)
	return string(jsonString)
}

type RequestLog struct {
	Code    int    `json:"code"`
	BaseURL string `json:"baseURL"`
	Path    string `json:"path"`
}

func (log *RequestLog) ToString() string {
	return fmt.Sprintf("%d | /%s", log.Code, log.Path)
}

func (log *RequestLog) ToJSON() string {
	jsonString, _ := json.Marshal(log)
	return string(jsonString)
}
