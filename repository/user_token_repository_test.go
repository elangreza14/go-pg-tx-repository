package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/elangreza14/go-pg-tx-repository/model"
	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v3"
)

func Test_userTokenRepository_CreateUserTX_Succeed(t *testing.T) {
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

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO users").
		WithArgs(req.ID, req.Username, req.Email, req.Password).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))
	mock.ExpectCommit()

	mockUserTokenRepo := NewUserTokenRepository(mock)
	if err := mockUserTokenRepo.CreateUserTX(context.Background(), req, func(res *model.User) error {
		return nil
	}); err != nil {
		t.Errorf("error was not expected while updating: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_userTokenRepository_CreateUserTX_Failed(t *testing.T) {
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

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO users").
		WithArgs(req.ID, req.Username, req.Email, req.Password).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))
	mock.ExpectRollback()

	mockUserTokenRepo := NewUserTokenRepository(mock)
	if err := mockUserTokenRepo.CreateUserTX(context.Background(), req, func(res *model.User) error {
		return errors.New("test err")
	}); err != nil {
		t.Logf("error was expected while updating: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
