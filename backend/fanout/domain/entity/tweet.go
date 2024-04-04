package entity

import (
	"github.com/mastar3104/twitter-clone/fanout/domain/value"
	"time"
)

type Tweet struct {
	tweetID   value.TweetID
	userID    value.UserID
	content   string
	tweetTime time.Time
}

// TweetID tweetIDのGetter
func (t Tweet) TweetID() value.TweetID {
	return t.tweetID
}

// UserID userIDのGetter
func (t Tweet) UserID() value.UserID {
	return t.userID
}

// TweetReconstructor repository等外部の値からEntityを再生成するために利用
func TweetReconstructor(
	tweetID value.TweetID,
	userID value.UserID,
	content string,
	tweetTime time.Time,
) Tweet {
	return Tweet{
		tweetID:   tweetID,
		userID:    userID,
		content:   content,
		tweetTime: tweetTime,
	}
}
