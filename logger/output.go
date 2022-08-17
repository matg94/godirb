package logger

import "fmt"

type RequestResult struct {
	Code    int    `json:"code"`
	BaseURL string `json:"url"`
	Path    string `json:"word"`
}

type Outputter struct {
	Results []RequestResult
}

func CreateOutput() *Outputter {
	return &Outputter{
		Results: []RequestResult{},
	}
}

func (out *Outputter) Log(Code int, BaseUrl, Path string) {
	out.Results = append(out.Results, RequestResult{
		Code:    Code,
		BaseURL: BaseUrl,
		Path:    Path,
	})
}

func (out *Outputter) GetStats() map[int]int {
	responses := map[int]int{}
	for _, res := range out.Results {
		if _, found := responses[res.Code]; found {
			responses[res.Code] += 1
		} else {
			responses[res.Code] = 0
		}
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
