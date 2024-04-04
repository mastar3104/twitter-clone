package repository

import (
	"github.com/mastar3104/twitter-clone/tweet/domain/entity"
	"github.com/mastar3104/twitter-clone/tweet/domain/value"
)

type UserRepository interface {
	Find(userID value.UserID) (entity.User, error)
}
