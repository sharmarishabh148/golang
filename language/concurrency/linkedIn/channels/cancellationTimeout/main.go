package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// After waits for the duration to elapse and then sends the current time
	// on the returned channel.
	timeout := time.After(2 * time.Second)
	done := make(chan bool)

	go func() {
		//time.Sleep(1001 * time.Millisecond)
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
		fmt.Printf("Function Took : %v", time.Since(t))
		done <- true
	}()

	select {
	case <-timeout:
		fmt.Printf("Timeout from time.After")
	case <-done:
		fmt.Printf("Function completed.")
		//default:
		//	fmt.Printf("Nothing")
	}

}
