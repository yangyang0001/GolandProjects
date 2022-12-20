package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCount struct {
	vmp map[string]int `json:"vmp"`
	mux sync.Mutex     `json:"mux"`
}

func main() {

	safe := SafeCount{vmp: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go safe.Inc("zhangsan")
	}

	time.Sleep(time.Second)
	fmt.Printf("last result = %v \n", safe.vmp["zhangsan"])

}

func (s *SafeCount) Inc(key string) int {

	s.mux.Lock()

	s.vmp[key]++

	defer s.mux.Unlock()

	return s.vmp[key]

}
