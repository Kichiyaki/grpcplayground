package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	"go.uber.org/zap"

	"github.com/Kichiyaki/grpcplayground/pb"

	"github.com/Kichiyaki/grpcplayground/cmd/internal"
	internaldomain "github.com/Kichiyaki/grpcplayground/internal"
	internalgrpc "github.com/Kichiyaki/grpcplayground/internal/grpc"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
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
		logger:  logger,
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
	logger  *zap.Logger
}

func newServer(cfg serverConfig) (*server, error) {
	lis, err := net.Listen("tcp", cfg.address)
	if err != nil {
		return nil, internaldomain.Wrap(err, "net.Listen")
	}

	grpc_zap.ReplaceGrpcLoggerV2(cfg.logger)

	srv := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(createPanicHandler(cfg.logger))),
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(cfg.logger),
		),
	)

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

func createPanicHandler(logger *zap.Logger) grpc_recovery.RecoveryHandlerFunc {
	return func(p interface{}) error {
		logger.Panic(fmt.Sprintf("%v", p), zap.Stack("stack"))
		return status.Errorf(codes.Internal, "internal server error")
	}
}
