package sql_manger

import "context"

type SQLManger interface {
	GetUserAccountByPhone(ctx context.Context, phoneNumber string) error
}
