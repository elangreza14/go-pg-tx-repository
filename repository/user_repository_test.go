package repository

import (
	"context"
	"testing"

	"github.com/elangreza14/go-pg-tx-repository/model"
	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v3"
)

func Test_userRepository_Create(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	req := &model.User{
		ID:       uuid.New(),
		Username: "a",
		Email:    "a",
		Password: []byte("a"),
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(req.ID, req.Username, req.Email, req.Password).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	mockUserRepo := NewUserRepository(mock)
	if err := mockUserRepo.Create(context.Background(), *req); err != nil {
		t.Errorf("error was not expected while updating: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
