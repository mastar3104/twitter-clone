package entity

import (
	"github.com/mastar3104/twitter-clone/tweet/domain/event"
	"github.com/mastar3104/twitter-clone/tweet/domain/value"
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

// Content contentのGetter
func (t Tweet) Content() string {
	return t.content
}

// TweetTime tweetTimeのGetter
func (t Tweet) TweetTime() time.Time {
	return t.tweetTime
}

// CreatTweet 投稿用ツイートの生成
func CreatTweet(user User, content string) (Tweet, *event.TweetPost) {
	return Tweet{
		tweetID:   value.CreateTweetID(),
		userID:    user.userID,
		content:   content,
		tweetTime: time.Now(),
	}, &event.TweetPost{}
}
