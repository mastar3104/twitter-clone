package repository

import (
	"github.com/mastar3104/twitter-clone/user/domain/entity"
	"github.com/mastar3104/twitter-clone/user/domain/value"
)

type UserRepository interface {
	Save(user entity.User)
	Find(userID value.UserID) (entity.User, error)
	Get() []entity.User
}
