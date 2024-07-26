package main

import (
	"sync"
)

type State struct {
	mu    sync.Mutex
	count int
}

func (s *State) setState(val int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.count = val

}

func main() {
}
