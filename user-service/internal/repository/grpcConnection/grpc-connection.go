package grpcConnection

import (
	"context"
	"user-service/proto/userAuth"
)

type GRPCConnections interface {
	GetUserAuthGRPcConnection(ctx context.Context) (userAuth.UserAuthClient, error)
}
