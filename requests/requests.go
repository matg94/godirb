package requests

import (
	"fmt"
	"net/http"

	"github.com/matg94/godirb/config"
	"github.com/matg94/godirb/data"
)

type RequestGenerator struct {
	BaseURL string
	Headers []config.HeaderConfig
	Cookie  string
}

func (rm *RequestGenerator) GetNextRequest(queue *data.WordQueue) (Request, error) {
	nextWord, err := queue.Next()
	if err == data.ErrEmptyQueue {
		return Request{}, err
	}
	return CreateRequest(rm.BaseURL, nextWord, rm.Cookie, rm.Headers), nil
}

type Request struct {
	URL     string
	Path    string
	Headers []config.HeaderConfig
	Cookie  string
}

func (r *Request) Send() (int, error) {
	fullURL := fmt.Sprintf("%s/%s", r.URL, r.Path)
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return -1, err
	}
	if len(r.Headers) != 0 {
		for _, header := range r.Headers {
			req.Header.Set(header.Header, header.Content)
		}
	}
	if r.Cookie != "" {
		cookie := http.Cookie{
			Name:  "cookie",
			Value: r.Cookie,
		}
		req.AddCookie(&cookie)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	return resp.StatusCode, nil
}

func CreateRequest(URL, path, cookie string, headers []config.HeaderConfig) Request {
	return Request{
		URL:     URL,
		Path:    path,
		Cookie:  cookie,
		Headers: headers,
	}
}
