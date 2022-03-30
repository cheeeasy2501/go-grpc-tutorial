package adder

import (
	"context"
)

// GRPCServer ...
type GRPCServer struct {
	UnimplementedAdderServer
}

func (s *GRPCServer) mustEmbedUnimplementedAdderServer() {}

// Add ...
func (s *GRPCServer) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	//делаем что-то и возвращаем message, описанный в proto
	return &AddResponse{Result: req.GetX() + req.GetY()}, nil
}
