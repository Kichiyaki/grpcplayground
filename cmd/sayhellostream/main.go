package main

import (
	"context"
	"io"
	"log"

	"go.uber.org/zap"

	"github.com/Kichiyaki/grpcplayground/pb"

	"github.com/Kichiyaki/grpcplayground/cmd/internal"
)

var names = [...]string{
	"John",
	"Winston",
	"Tommy",
}

func main() {
	logger, err := internal.NewLogger()
	if err != nil {
		log.Fatalln("internal.NewLogger:", err)
	}
	defer func() {
		_ = logger.Sync()
	}()

	conn, err := internal.NewLocalGRPCClientConn()
	if err != nil {
		logger.Fatal("grpc.Dial: " + err.Error())
	}
	defer func() {
		_ = conn.Close()
	}()

	client := pb.NewPlaygroundClient(conn)

	stream, err := client.SayHelloStream(context.Background())
	if err != nil {
		logger.Fatal("client.SayHelloStream: " + err.Error())
	}

	waitc := make(chan struct{}, 1)
	defer close(waitc)

	go func(logger *zap.Logger, stream pb.Playground_SayHelloStreamClient, done chan<- struct{}) {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				done <- struct{}{}
				return
			}
			if err != nil {
				logger.Fatal("stream.Recv: " + err.Error())
			}

			logger.Info(in.GetMessage())
		}
	}(logger, stream, waitc)

	for _, name := range names {
		if err := stream.Send(&pb.SayHelloRequest{Name: name}); err != nil {
			logger.Fatal("stream.Send: "+err.Error(), zap.String("name", name))
		}
	}

	if err := stream.CloseSend(); err != nil {
		logger.Fatal("stream.Send: " + err.Error())
	}

	<-waitc
}
