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

func UserFromId(idString string) (*User, error) {
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

type GetUserDTO struct {
	UUID string `json:"uuid"`
}
