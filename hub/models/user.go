package models

import "github.com/google/uuid"

type User struct {
	UUID uuid.UUID
}

func NewUser(idString string) (*User, error) {
	id, err := uuid.Parse(idString)
	if err != nil {
		return nil, err
	}
	return &User{
		UUID: id,
	}, nil
}

type RegisterUserDTO struct {
	UUID string `json:"uuid"`
}
