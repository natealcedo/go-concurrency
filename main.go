package main

import (
	"fmt"
)

func main() {
	// There are 2 types of channels: buffered and unbuffered
	// An unbuffered channel is always going to block regardless of whether its full or not
	resultch := make(chan string, 1) // unbuffered channel

	resultch <- "foo" // -> is now full and will block

	result := <-resultch
	fmt.Println(result)

}
