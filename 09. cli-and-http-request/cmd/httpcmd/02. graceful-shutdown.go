package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"sync"
)

func startHttpServer(wg *sync.WaitGroup) *http.Server {
	srv := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})

	go func() {
		defer wg.Done() // let main know we are done cleaning up

		// always returns error. ErrServerClosed on graceful close
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	http.HandleFunc("/shutdown-server", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("main: stopping HTTP server")
		io.WriteString(w, "Request to shutdown server.\n")

		io.WriteString(w, "Shutting down the server\n")
		if err := srv.Shutdown(context.TODO()); err != nil {
			panic(err) // failure/timeout shutting down the server gracefully
		}
	})

	// returning reference so caller can call Shutdown()
	return srv
}

func main() {
	log.Printf("main: starting HTTP server")
	httpServerExitDone := &sync.WaitGroup{}
	httpServerExitDone.Add(1)
	startHttpServer(httpServerExitDone)
	log.Printf("main: serving for 10 seconds")
	// wait for goroutine started in startHttpServer() to stop
	httpServerExitDone.Wait()
	log.Printf("main: done. exiting")
}
