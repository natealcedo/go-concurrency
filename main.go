package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "username", "natealcedo")

	val, err := fetchUserId(ctx)

	if err != nil {
		_ = fmt.Errorf("error fetching value: %v", err)
		log.Fatal(err)
	}

	fmt.Printf("The response is -> %s and took %+v\n", val, time.Since(start))
}

func fetchUserId(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	username := ctx.Value("username")
	fmt.Println(username)

	type result struct {
		userId string
		err    error
	}

	resCh := make(chan result, 5)

	go func() {
		userId, err := thirdPartyFetchValue()
		resCh <- result{userId, err}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resCh:
		return res.userId, res.err
	}
}

func thirdPartyFetchValue() (string, error) {
	time.Sleep(time.Millisecond * 90)
	return "user id 1", nil
}
