package grpcConnection

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"user-service/proto/userAuth"
)

var etcdConn *grpc.ClientConn

type GRPCConnectionsImpl struct {
}

func (r *GRPCConnectionsImpl) GetUserAuthGRPcConnection(ctx context.Context) (userAuth.UserAuthClient, error) {
	etcdConn, err := grpc.DialContext(context.Background(), "user-auth-grpc:10021", grpc.WithInsecure())
	if err != nil {
		log.Error().Any("Failed to dial gRPC server: %v", err)
		return nil, err
	}
	client := userAuth.NewUserAuthClient(etcdConn)
	return client, nil
}
