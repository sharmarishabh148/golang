package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// retryTimeout: You need to validate if something can be done with no error
// but it may take time before this is true. You set a retry interval to create
// a delay before you retry the call and you use the context to set a timeout.
func retryTimeout(ctx context.Context,
	retryInterval time.Duration,
	check func(ctx context.Context) error) {

	for {
		fmt.Println("perform user check call")
		if err := check(ctx); err == nil {
			fmt.Println("work finished successfully")
			return
		}

		fmt.Println("check if timeout has expired")
		if ctx.Err() != nil {
			fmt.Println("time expired 1 :", ctx.Err())
			return
		}

		fmt.Printf("wait %s before trying again\n", retryInterval)
		t := time.NewTimer(retryInterval) //time.After also can be used

		select {
		case <-ctx.Done():
			fmt.Println("timed expired 2 :", ctx.Err())
			t.Stop()
			return
		case <-t.C:
			fmt.Println("retry again")
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	retryTimeout(ctx, time.Second, func(ctx context.Context) error { return errors.New("always fail") })
}
