package repository

import (
	"context"
	"testing"

	"github.com/elangreza14/go-pg-tx-repository/model"
	"github.com/pashagolub/pgxmock/v3"
)

func Test_tokenRepository_Create(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	req := model.Token{}

	mock.ExpectExec("INSERT INTO tokens").
		WithArgs(req.ID, req.UserID, req.Token, req.TokenType, req.IP, req.IssuedAt, req.ExpiredAt, req.Duration).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	mockTokenRepo := NewTokenRepository(mock)
	if err := mockTokenRepo.Create(context.Background(), req); err != nil {
		t.Errorf("error was not expected while updating: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
