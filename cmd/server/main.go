package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/Kichiyaki/grpcplayground/pb"

	"github.com/Kichiyaki/grpcplayground/cmd/internal"
	internaldomain "github.com/Kichiyaki/grpcplayground/internal"
	internalgrpc "github.com/Kichiyaki/grpcplayground/internal/grpc"
	"google.golang.org/grpc"
)

func main() {
	logger, err := internal.NewLogger()
	if err != nil {
		log.Fatalln("internal.NewLogger:", err)
	}
	defer func() {
		_ = logger.Sync()
	}()

	srv, err := newServer(serverConfig{
		address: "localhost:" + internal.GetPort(),
	})
	if err != nil {
		logger.Fatal("newServer: " + err.Error())
	}

	go func(srv *server, logger *zap.Logger) {
		ctx, stop := signal.NotifyContext(
			context.Background(),
			os.Interrupt,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)
		defer stop()

		<-ctx.Done()

		logger.Info("shutdown signal received")

		srv.GracefulStop()
	}(srv, logger)

	logger.Info(
		"listening and serving",
		zap.String("address", srv.lis.Addr().String()),
	)

	if err := srv.Serve(); err != nil {
		logger.Fatal("srv.Serve: " + err.Error())
	}

	logger.Info("shutdown completed")
}

type server struct {
	lis     net.Listener
	grpcSrv *grpc.Server
}

type serverConfig struct {
	address string
}

func newServer(cfg serverConfig) (*server, error) {
	lis, err := net.Listen("tcp", cfg.address)
	if err != nil {
		return nil, internaldomain.Wrap(err, "net.Listen")
	}

	srv := grpc.NewServer()

	pb.RegisterPlaygroundServer(srv, internalgrpc.NewPlaygroundServer())

	return &server{
		lis:     lis,
		grpcSrv: srv,
	}, nil
}

func (s *server) Serve() error {
	return s.grpcSrv.Serve(s.lis)
}

func (s *server) GracefulStop() {
	s.grpcSrv.GracefulStop()
}
