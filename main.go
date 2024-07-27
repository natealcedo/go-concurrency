package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	ID       int      `json:"id"`
	Comments []string `json:"comments"`
	Likes    int      `json:"likes"`
	Friends  []int    `json:"friends"`
}

type Response struct {
	data any
	err  error
}

func handleGetUserProfile(id int) (*UserProfile, error) {
	responseCh := make(chan *Response, 3)
	wg := &sync.WaitGroup{}
	go getComments(id, responseCh, wg)
	go getLikes(id, responseCh, wg)
	go getFriends(id, responseCh, wg)

	wg.Add(3)
	wg.Wait()
	close(responseCh)

	userProfile := &UserProfile{ID: id}
	// Keep ranging. But when to stop?
	for resp := range responseCh {
		if resp.err != nil {
			return nil, resp.err
		}
		switch msg := resp.data.(type) {
		case []string:
			userProfile.Comments = msg
		case int:
			userProfile.Likes = msg
		case []int:
			userProfile.Friends = msg
		}
	}

	return userProfile, nil
}

func getComments(id int, responseCh chan *Response, wg *sync.WaitGroup) {
	time.Sleep(200 * time.Millisecond)
	comments := []string{
		"yeah buddy",
		"I didn't know that",
	}
	responseCh <- &Response{data: comments, err: nil}
	wg.Done()
}

func getLikes(id int, responseCh chan *Response, wg *sync.WaitGroup) {
	time.Sleep(200 * time.Millisecond)
	responseCh <- &Response{data: 33, err: nil}
	wg.Done()
}

func getFriends(id int, responseCh chan *Response, wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	responseCh <- &Response{data: []int{11, 12, 123, 6456}, err: nil}
	wg.Done()
}

func main() {
	start := time.Now()
	userProfile, err := handleGetUserProfile(10)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User Profile: %+v\n", userProfile)
	fmt.Println("Time taken: ", time.Since(start))
}
