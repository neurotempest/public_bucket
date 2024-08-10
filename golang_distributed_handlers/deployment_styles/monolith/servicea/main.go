package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

var apiHost = flag.String("api_host", "", "host for REST api server")
var apiPort = flag.String("api_port", "1234", "port for REST api server")

func runWithListener(ctx context.Context, listener net.Listener) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()


	srv := NewServer(
		// TODO add server dependencies
	)

	httpServer := &http.Server{
		Handler: srv,
	}
	go func() {
		if err := httpServer.Serve(listener); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		// make a new context for the Shutdown (thanks Alessandro Rosetti)
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(ctx, 10 * time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}

func run(ctx context.Context, host string, port string) error {
	addr := net.JoinHostPort(host, port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("error creating listener: %s\n", err)
	}
	log.Printf("listening on %s\n", listener.Addr())
	return runWithListener(ctx, listener)
}

func main() {
	ctx := context.Background()
	if err := run(ctx, *apiHost, *apiPort); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
