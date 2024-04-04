package repository

import (
	"github.com/mastar3104/twitter-clone/tweet/domain/entity"
	"github.com/mastar3104/twitter-clone/tweet/domain/event"
)

type TweetRepository interface {
	Save(tweet entity.Tweet, event event.TweetPost)
}
