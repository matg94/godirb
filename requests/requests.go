package requests

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/matg94/godirb/config"
	"github.com/matg94/godirb/data"
)

func GetNextRequest(queue *data.WordQueue, baseURL string, conf config.RequestConfig) (*Request, error) {
	nextWord, err := queue.Next()
	if err == data.ErrEmptyQueue {
		return &Request{}, err
	}
	return CreateRequest(baseURL, nextWord, conf), nil
}

type Request struct {
	URL     string
	Path    string
	Headers []config.HeaderConfig
	Cookie  string
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

func CreateRequest(baseURL, word string, conf config.RequestConfig) *Request {
	return &Request{
		URL:     baseURL,
		Path:    word,
		Cookie:  conf.Cookie,
		Headers: conf.Headers,
	}
}
