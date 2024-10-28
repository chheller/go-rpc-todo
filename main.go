package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/chheller/go-rpc-todo/config"
	"github.com/chheller/go-rpc-todo/server"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

func main() {
	env := config.GetEnvironment();
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(env.ApplicationConfiguration.LogLevel)

	srv := grpc.NewServer();

	// Register services with the server
	router := server.RPCServer{Server: srv}
	router.Init()

	// Create a channel to recieve shutdown signals.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)

	// Run the server in a never ending goroutine.
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", env.ApplicationConfiguration.Port))
		if err != nil {
			log.Panicf("Error stopping server %s", err)
		}
		log.Printf("Server listening on %s", lis.Addr().String())
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}

	}()

	<-stop
	log.Printf("Shutting down server")
	srv.Stop()
	log.Println("Server gracefully stopped")
}
