package redis_manger

import "context"

type RedisManger interface {
	SetUserOTPRedisModel(ctx context.Context, pinToken string) error
	GetUserOTPRedisModel(ctx context.Context, pinToken string) error
}
