package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/chheller/go-rpc-todo/config"
	"github.com/chheller/go-rpc-todo/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	log "github.com/sirupsen/logrus"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	return token == "some-secret-token"
}

func unaryInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// authentication (token verification)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}
	m, err := handler(ctx, req)
	if err != nil {
		log.Printf("RPC failed with error: %v", err)
	}
	return m, err
}

func main() {
	env := config.GetEnvironment();
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(env.ApplicationConfiguration.LogLevel)

	// TODO: Load from environment
	// TODO: Figure out how to get certs into container from k8s
	
	creds, err := credentials.NewServerTLSFromFile(env.ApplicationConfiguration.HttpsCertificatePath, env.ApplicationConfiguration.HttpsKeyPath)
	if err != nil {
		log.Fatalf("failed to create credentials: %v", err)
	}

	srv := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(unaryInterceptor));

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
