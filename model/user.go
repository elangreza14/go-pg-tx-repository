package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Email    string    `db:"email"`
	Password []byte    `db:"password"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func NewUser(userName, email, password string) (*User, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       id,
		Username: userName,
		Email:    email,
		Password: pass,
	}, nil
}

func (u *User) IsPasswordValid(reqPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(reqPassword))
	return err == nil
}

func (u User) TableName() string {
	return "users"
}

func (u User) Columns() []string {
	return []string{
		"id",
		"username",
		"email",
		"password",
	}
}

// to create in DB
func (u User) Data() map[string]any {
	return map[string]any{
		"id":       u.ID,
		"username": u.Username,
		"email":    u.Email,
		"password": u.Password,
	}
}
