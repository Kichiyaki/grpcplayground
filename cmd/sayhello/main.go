package main

import (
	"context"
	"flag"
	"log"

	"github.com/Kichiyaki/grpcplayground/pb"

	"github.com/Kichiyaki/grpcplayground/cmd/internal"
)

func main() {
	logger, err := internal.NewLogger()
	if err != nil {
		log.Fatalln("internal.NewLogger:", err)
	}
	defer func() {
		_ = logger.Sync()
	}()

	name := ""

	flag.StringVar(&name, "name", "anonymous", "who should the server say hello to")
	flag.Parse()

	conn, err := internal.NewLocalGRPCClientConn()
	if err != nil {
		logger.Fatal("grpc.Dial: " + err.Error())
	}
	defer func() {
		_ = conn.Close()
	}()

	client := pb.NewPlaygroundClient(conn)

	resp, err := client.SayHello(context.Background(), &pb.SayHelloRequest{Name: name})
	if err != nil {
		logger.Fatal("client.SayHello: " + err.Error())
	}

	logger.Info(resp.GetMessage())
}
