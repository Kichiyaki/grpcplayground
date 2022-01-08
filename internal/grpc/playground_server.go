package grpc

import (
	"context"
	"io"

	"github.com/Kichiyaki/grpcplayground/pb"
)

type PlaygroundServer struct {
	pb.UnimplementedPlaygroundServer
}

func NewPlaygroundServer() *PlaygroundServer {
	return &PlaygroundServer{}
}

func (srv *PlaygroundServer) SayHello(_ context.Context, r *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{Message: newHelloMessage(r.GetName())}, nil
}

func (srv *PlaygroundServer) SayHelloStream(stream pb.Playground_SayHelloStreamServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		if err := stream.Send(&pb.SayHelloResponse{Message: newHelloMessage(in.GetName())}); err != nil {
			return err
		}
	}
}

func newHelloMessage(name string) string {
	return "Hello " + name + "!"
}
