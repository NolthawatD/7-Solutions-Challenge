package main

import (
	"client/services"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	creds := insecure.NewCredentials()
	cc, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	countingClient := services.NewCountingClient(cc)
	countingServices := services.NewCountingService(countingClient)

	// err = countingServices.Hello("Nolthawat")
	err = countingServices.CountBeef()
	if err != nil {
		if grpcErr, ok := status.FromError(err); ok {
			log.Printf("[%v] %v", grpcErr.Code(), grpcErr.Message())
		} else {
			log.Fatal(err)
		}
	}

}
