package model

import (
	"errors"
	"time"

	"github.com/koba1108/go-mongodb/internals/helper"
)

var (
	ErrUserNameEmpty  = errors.New("user name is empty")
	ErrUserEmailEmpty = errors.New("user email is empty")
)

type User struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func NewUser(name, email string) (*User, error) {
	if name == "" {
		return nil, ErrUserNameEmpty
	}
	if email == "" {
		return nil, ErrUserEmailEmpty
	}
	return &User{
		ID:        helper.NewULID().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}, nil
}
