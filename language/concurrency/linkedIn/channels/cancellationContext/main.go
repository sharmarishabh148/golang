package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	duration := 1500 * time.Millisecond
	//background is  Empty context, since we dont have a already existing
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	//if forgetting memory leaks
	defer cancel()

	// If unbuffered very careful, because when timer fires we will be out of select call
	// and there will be noone to recieve on that channel
	ch := make(chan string, 1)

	go func() {
		// Lets say some work we are doing some call to db or http grpc request
		t := time.Now()
		resp, err := http.Get("https://httpbin.org/get")
		if err != nil {
			fmt.Print(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(string(body))
		fmt.Printf("Function Took : %v\n", time.Since(t))
		// signaling data, sender is done , reciever can recieve whenever it want
		ch <- "i am done"
	}()

	select {
	case d := <-ch:
		fmt.Printf("work completed : signal recieved %s", d)
	case <-ctx.Done():
		fmt.Printf("Timeout from context ")
	}

}
