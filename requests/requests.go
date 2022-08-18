package requests

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/matg94/godirb/config"
	"github.com/matg94/godirb/data"
)

type RequestGenerator struct {
	BaseURL string
	Headers []config.HeaderConfig
	Cookie  string
	Client  *http.Client
}

func (rm *RequestGenerator) GetNextRequest(queue *data.WordQueue) (*Request, error) {
	nextWord, err := queue.Next()
	if err == data.ErrEmptyQueue {
		return &Request{}, err
	}
	return CreateRequest(rm.BaseURL, nextWord, rm.Cookie, rm.Headers), nil
}

type Request struct {
	URL     string
	Path    string
	Headers []config.HeaderConfig
	Cookie  string
}

type Timer struct {
	startTime time.Time
}

func (t *Timer) Time(code string) {
	fmt.Println("Time: ", time.Since(t.startTime), " Code: ", code)
}

func (r *Request) Send(client *http.Client) (int, error) {
	fullURL := fmt.Sprintf("%s/%s", r.URL, r.Path)
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return -1, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	if _, err := io.Copy(ioutil.Discard, resp.Body); err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}

func CreateRequest(URL, path, cookie string, headers []config.HeaderConfig) *Request {
	return &Request{
		URL:     URL,
		Path:    path,
		Cookie:  cookie,
		Headers: headers,
	}
}
