package data

import (
	"errors"
	"sync"
)

type WordQueue struct {
	mut   sync.Mutex
	words []string
}

var ErrEmptyQueue = errors.New("empty queue")

func (w *WordQueue) Next() (string, error) {
	w.mut.Lock()
	defer w.mut.Unlock()

	if len(w.words) == 0 {
		return "", ErrEmptyQueue
	}
	next := w.words[0]
	w.words = w.words[1:]
	return next, nil
}

func (w *WordQueue) AddList(words []string) {
	w.words = append(w.words, words...)
}

func (w *WordQueue) GetAll() []string {
	return w.words
}

func CreateWordQueue() *WordQueue {
	return &WordQueue{
		words: []string{},
	}
}
