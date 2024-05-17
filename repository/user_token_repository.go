package repository

import (
	"context"

	"github.com/elangreza14/go-pg-tx-repository/model"
)

type userTokenRepository struct {
	txRepo *PostgresTransactionRepo
}

func NewUserTokenRepository(
	tx PgxTXer,
) *userTokenRepository {
	return &userTokenRepository{
		txRepo: NewPostgresTransactionRepo(tx),
	}
}

func (ur *userTokenRepository) CreateUserTX(ctx context.Context, req *model.User, callback func(*model.User) error) error {
	return ur.txRepo.WithTX(ctx, func(tx QueryPgx) error {
		userRepo := NewUserRepository(tx)
		var err error
		if err = userRepo.Create(ctx, *req); err != nil {
			return err
		}

		req.Username = "update a"

		// if err := userRepo.Edit(ctx, *req); err != nil {
		// 	return err
		// }

		return callback(req)
	})
}
