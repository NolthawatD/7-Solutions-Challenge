package services

import (
	context "context"
	"fmt"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type countingServer struct{}

func NewCountingServer() CountingServer {
	return countingServer{}
}

func (countingServer) mustEmbedUnimplementedCountingServer() {}

func (countingServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {

	if req.Name == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Name is required",
		)
	}

	result := fmt.Sprintf("Hello %v", req.Name)
	res := HelloResponse{
		Result: result,
	}
	return &res, nil
}
