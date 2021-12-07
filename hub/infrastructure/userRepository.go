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

func (u *UserRepository) GetUsers(ctx echo.Context, ownId string) ([]string, error) {
	filterFunc := func(ids []string) []string {
		var filteredIds []string
		for _, v := range ids {
			if v != ownId {
				filteredIds = append(filteredIds, v)
			}
		}
		return filteredIds
	}
	userIds, err := u.cache.GetAll(ctx, filterFunc)
	if err != nil {
		return nil, err
	}
	return userIds, nil
}
