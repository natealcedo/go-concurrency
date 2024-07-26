package main

import (
	"fmt"
	"time"
)

func main() {
	res := fetchResource(10)
	fmt.Println(res)
}

func fetchResource(val int) string {
	time.Sleep(2 * time.Second)

	return fmt.Sprintf("result %d", val)
}
