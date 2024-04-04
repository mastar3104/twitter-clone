package event

import (
	"github.com/mastar3104/twitter-clone/user/domain/value"
)

// UserFollow ユーザがフォローした際に発生するイベントのトリガー
type UserFollow struct{}

func (t UserFollow) Publish(userID value.UserID, followUserID value.UserID, publisher TweetPostPublisher) {
	publisher.Send(userID, followUserID)
}

type TweetPostPublisher interface {
	Send(userID value.UserID, followUserID value.UserID)
}
