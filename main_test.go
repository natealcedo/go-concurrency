package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestState(t *testing.T) {
	state := &State{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(count int) {
			state.setState(count + 1)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("%+v\n", state)
}
