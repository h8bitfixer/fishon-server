package grpcConnection

import (
	"admin-user-service/proto/userAuth"
	"context"
)

type GRPCConnections interface {
	GetUserAuthGRPcConnection(ctx context.Context) (userAuth.UserAuthClient, error)
}
