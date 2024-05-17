package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/pashagolub/pgxmock/v3"
)

func TestPostgresTransactionRepo_WithTX_Succeed(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	mock.ExpectBegin()
	mock.ExpectCommit()

	mockUserRepo := NewPostgresTransactionRepo(mock)
	err = mockUserRepo.WithTX(context.Background(), func(tx QueryPgx) error {
		return nil
	})

	if err != nil {
		t.Errorf("error was not expected while updating: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostgresTransactionRepo_WithTX_Failed(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	mock.ExpectBegin()
	mock.ExpectRollback()

	mockUserRepo := NewPostgresTransactionRepo(mock)
	err = mockUserRepo.WithTX(context.Background(), func(tx QueryPgx) error {
		return errors.New("test err")
	})

	if err != nil {
		t.Logf("error was not expected while updating: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
