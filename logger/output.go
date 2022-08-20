package logger

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type RequestResult struct {
	Code    int    `json:"code"`
	BaseURL string `json:"url"`
	Path    string `json:"word"`
	Ok      bool   `json:"ok"`
}

func (r *RequestResult) toString() string {
	return fmt.Sprintf("%d | /%s", r.Code, r.Path)
}

type Outputter struct {
	Results     []RequestResult
	DisplayLive bool
	mut         sync.Mutex
}

func CreateOutput(displayLive bool) *Outputter {
	return &Outputter{
		Results:     []RequestResult{},
		DisplayLive: displayLive,
	}
}

func (out *Outputter) Log(Code int, BaseUrl, Path string, ok bool) {
	out.mut.Lock()
	defer out.mut.Unlock()
	result := RequestResult{
		Code:    Code,
		BaseURL: BaseUrl,
		Path:    Path,
		Ok:      ok,
	}
	out.Results = append(out.Results, result)
	if out.DisplayLive && ok {
		fmt.Println(result.toString())
	}
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
	fmt.Println("Total Requests Sent: ", len(out.Results))
	for v, k := range responses {
		fmt.Printf("%d : %d\n", v, k)
	}
}

func (out *Outputter) OutputFile(file string) error {
	if file != "" {
		f, err := os.Create(file)
		if err != nil {
			return err
		}
		defer f.Close()
		var output string
		for _, lg := range out.Results {
			output += lg.toString() + "\n"
		}
		ioutil.WriteFile(file, []byte(output), 0644)
	} else {
		return errors.New("output file was not given")
	}
	return nil
}
