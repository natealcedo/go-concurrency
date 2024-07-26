package main

import "fmt"

func main() {
	msgCh := make(chan string, 128)

	msgCh <- "A"
	msgCh <- "B"
	msgCh <- "C"
	close(msgCh)

	for {
		msg, ok := <-msgCh
		if !ok {
			break
		}
		fmt.Println(msg)
	}

	fmt.Println("Done reading all the messages")
}
