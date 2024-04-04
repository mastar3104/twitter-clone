package event

import (
	"github.com/mastar3104/twitter-clone/tweet/domain/value"
)

// TweetPost ユーザがツイートした際に発生するイベントのトリガー
type TweetPost struct{}

func (t TweetPost) Publish(userID value.UserID, tweetID value.TweetID, publisher TweetPostPublisher) {
	publisher.Send(userID, tweetID)
}

type TweetPostPublisher interface {
	Send(userID value.UserID, tweetID value.TweetID)
}
