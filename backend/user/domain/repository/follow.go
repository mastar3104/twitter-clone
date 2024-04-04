package repository

import (
	"github.com/mastar3104/twitter-clone/user/domain/entity"
	"github.com/mastar3104/twitter-clone/user/domain/event"
	"github.com/mastar3104/twitter-clone/user/domain/value"
)

type FollowRepository interface {
	Find(userID value.UserID, followUserID value.UserID) (entity.Follow, error)
	Save(follow entity.Follow, event event.UserFollow)
}
