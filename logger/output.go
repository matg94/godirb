package logger

import (
	"fmt"
	"sync"
)

type RequestResult struct {
	Code    int    `json:"code"`
	BaseURL string `json:"url"`
	Path    string `json:"word"`
}

type Outputter struct {
	Results []RequestResult
	mut     sync.Mutex
}

func CreateOutput() *Outputter {
	return &Outputter{
		Results: []RequestResult{},
	}
}

func (out *Outputter) Log(Code int, BaseUrl, Path string) {
	out.mut.Lock()
	defer out.mut.Unlock()
	out.Results = append(out.Results, RequestResult{
		Code:    Code,
		BaseURL: BaseUrl,
		Path:    Path,
	})
}

func (out *Outputter) GetStats() map[int]int {
	responses := map[int]int{}
	for _, res := range out.Results {
		responses[res.Code] += 1
	}
	return responses
}

func (out *Outputter) OutputPretty() {
	responses := out.GetStats()
	fmt.Println("Total requests sent: ", len(out.Results))
	for v, k := range responses {
		fmt.Printf("%d : %d\n", v, k)
	}
}
