package internal

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewLocalGRPCClientConn() (*grpc.ClientConn, error) {
	return grpc.Dial(
		"localhost:"+GetPort(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
}
