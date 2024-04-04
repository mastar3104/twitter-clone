package model

import (
	"github.com/mastar3104/twitter-clone/tweet/domain/value"
	"time"
)

// Tweet 画面に表示したいツイート情報のモデル
type Tweet struct {
	UserId    value.UserID
	UserName  string
	TweetId   value.TweetID
	Content   string
	TweetTime time.Time
}
