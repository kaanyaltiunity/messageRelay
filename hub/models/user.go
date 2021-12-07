package models

import "github.com/google/uuid"

type User struct {
	UUID uuid.UUID
}

func NewUser() (*User, error) {
	return &User{
		UUID: uuid.New(),
	}, nil
}

type RegisterUserDTO struct {
	UUID string `json:"uuid"`
}
