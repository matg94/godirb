package logger

import (
	"fmt"
	"os"
	"sync"
	"text/tabwriter"
)

type RequestCounterMap struct {
	mut sync.Mutex
	Map map[int]int
}

func CreateRequestCounterMap() *RequestCounterMap {
	return &RequestCounterMap{
		mut: sync.Mutex{},
		Map: map[int]int{},
	}
}

func (r *RequestCounterMap) Add(key int) {
	r.mut.Lock()
	defer r.mut.Unlock()

	if _, ok := r.Map[key]; !ok {
		r.Map[key] = 0
	}
	r.Map[key] += 1
}

func (r *RequestCounterMap) Print() {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "Code\t|\tCount")
	fmt.Fprintln(w, "-----\t|\t-----")
	for k, v := range r.Map {
		fmt.Fprintf(w, "%d\t|\t%d\n", k, v)
	}
	w.Flush()
}
