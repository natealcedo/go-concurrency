package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitCh chan any
	msgCh  chan string
}

func newServer() *Server {
	return &Server{
		quitCh: make(chan any),
		msgCh:  make(chan string, 128),
	}
}

func (s *Server) quit() {
	s.quitCh <- struct{}{}

}

func (s *Server) start() {
	fmt.Println("Server Starting")
	s.loop() // This is going to block
}

func (s *Server) sendMsg(msg string) {
	s.msgCh <- msg
}

func (s *Server) loop() {
mainLoop:
	for {
		select {
		case <-s.quitCh:
			fmt.Println("Quitting Server")
			break mainLoop
		case msg := <-s.msgCh:
			s.handleMsg(msg)
		}
	}
	fmt.Println("Server is shutting down gracefully")
}

func (s *Server) handleMsg(msg string) {
	fmt.Println("Received a message: ", msg)
}

func main() {
	s := newServer()
	go func() {
		time.Sleep(time.Second * 5)
		s.quit()
	}()
	s.start()
}
