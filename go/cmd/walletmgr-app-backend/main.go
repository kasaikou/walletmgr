package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

func main() {

	port := 15000
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	go func() {
		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()

	fmt.Printf("listen:grpc:%d", port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	server.GracefulStop()
}
