package grpc

import (
	"context"

	"github.com/Kichiyaki/grpcplayground/pb"
)

type PlaygroundServer struct {
	pb.UnimplementedPlaygroundServer
}

func NewPlaygroundServer() *PlaygroundServer {
	return &PlaygroundServer{}
}

func (srv *PlaygroundServer) SayHello(ctx context.Context, r *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{Message: "Hello " + r.GetName() + "!"}, nil
}
