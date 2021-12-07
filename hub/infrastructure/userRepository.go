package infrastructure

import (
	"messageHub/models"

	"github.com/labstack/echo"
)

type UserRepository struct {
	cache Cache
}

func NewUserRepository(cache Cache) *UserRepository {
	return &UserRepository{
		cache: cache,
	}
}

func (u *UserRepository) Register(ctx echo.Context, user *models.User) error {
	err := u.cache.Register(ctx, user.UUID.String())
	return err
}
