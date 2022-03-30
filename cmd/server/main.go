package main

import (
	"google.golang.org/grpc"
	"grpcadder/pkg/adder"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	grpcSrv := &adder.GRPCServer{}

	adder.RegisterAdderServer(s, grpcSrv)

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
