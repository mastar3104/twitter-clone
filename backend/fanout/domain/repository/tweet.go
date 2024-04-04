package repository

import (
	"github.com/mastar3104/twitter-clone/fanout/domain/entity"
	"github.com/mastar3104/twitter-clone/fanout/domain/value"
)

type TweetRepository interface {
	Get(userID value.UserID) []entity.Tweet
}
