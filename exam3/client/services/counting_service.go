package services

import (
	context "context"
	"fmt"
)

type CountingService interface {
	Hello(name string) error
}

type countingService struct {
	countingClient CountingClient
}

func NewCountingService(countingClient CountingClient) CountingService {
	return countingService{countingClient}
}

func (base countingService) Hello(name string) error {
	req := HelloRequest{
		Name: name,
	}

	res, err := base.countingClient.Hello(context.Background(), &req)
	if err != nil {
		return err
	}

	fmt.Printf("Service : Hello\n")
	fmt.Printf("Request : %v\n", req.Name)
	fmt.Printf("Response: %v\n", res.Result)

	return nil
}
