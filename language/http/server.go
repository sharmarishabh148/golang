package main

/*
You don't use Go routines primarily to take advantage of different CPU cores
(although this can also be a benefit). You use Go routines to keep your program
from blocking while doing something that takes a while.

In the case of a web application, most requests are usually not CPU intensive at
all. But they usually spend a lot of time (in computer terms) waiting around for
things to happen. They wait for DNS lookups on the request hostname, they wait for
the database to look up user credentials to establish a session, they wait for the
database to store or return rows to produce the HTTP response, etc.

Without go routines, while doing these things, your server would be unable to do
anything else. So if your typical HTTP request took, say, 1 second, to look up DNS,
validate an authorization cookie, look up results from a database, and send a response,
no other HTTP client could be served simultaneously.

Fortunately, the http package, which is used by Gorilla (and practically every other web
framework for Go) already uses Go routines to handle requests. So you're already using
(at least) one Go routine per HTTP request. Whether it makes sense to use additional go
routines is more up to your application design than "using more CPU cores."

Some suggestions:
If you're doing multiple things to serve a request that could be done in parallel,
use go routines. Suppose you need to do 3 DB queries, you can do each one in a Go
routine, so they run simultaneously. They may run on different CPU cores, but that's
completely irrelevant to actual performance, since each one is essentially doing nothing,
and just waiting for the database. But you'll still improve performance.
(For completeness sake,I'll mention that goroutines aren't the only way to run database
queries in parallel--but they are the idiomatic Go way, and the easy way in Go).

If you have tasks that are run by an HTTP request, but don't affect the HTTP response,
you can run them in a Go routine, so your response can return sooner. Logging is a common
example. You probably want to log your HTTP requests, but if the logger is slow, you probably
don't want to wait for logging to finish before sending your HTTP response. So do it in a go
routine. The logger can take it's sweet time and finish after the client has already received
their response. Another example is sending emails--please don't wait for an email to be sent
before responding to an HTTP client.

*/
import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func main() {
	var err error
	startAppServer()
}
func mapUrls() {
	router.HandleFunc("/", Index)
}
func startAppServer() {
	var wait time.Duration
	var waitDesc string = `the duration for which the server gracefully wait for existing
						connections to finish - e.g. 15s or 1m`
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, waitDesc)

	flag.Parse()

	fmt.Println("Server Starting...")
	mapUrls()

	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      router,
		WriteTimeout: 120000 * time.Millisecond,
		IdleTimeout:  60 * time.Second,
	}
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("Exiting Server...")
			//panic(err)
		}
	}()

	// Buffered channel , that means Blocking
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Many more pages to come"))
}
