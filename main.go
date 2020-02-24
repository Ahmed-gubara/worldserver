package main

import "net"

import "fmt"

import "google.golang.org/grpc"

type server struct {
}

const port int = 6000

func main() {
	lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))
	grpcServer := grpc.NewServer()
	grpcServer.Serve(lis)
}
