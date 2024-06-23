package main

import (
	"fmt"
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	port := ":50051"
	s := grpc.NewServer()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterCountingServer(s, services.NewCountingServer())
	reflection.Register(s)

	fmt.Printf("gRPC server listening on port %s", port)
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
