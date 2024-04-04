package repository

import (
	"github.com/mastar3104/twitter-clone/shared/infrastucture/adapter"
	"github.com/mastar3104/twitter-clone/tweet/domain/entity"
	"github.com/mastar3104/twitter-clone/tweet/domain/event"
	"github.com/mastar3104/twitter-clone/tweet/infrastructure/publisher"
)

type TweetRepository struct {
	Handler *adapter.MySQLHandler
}

func (t TweetRepository) Save(tweet entity.Tweet, event event.TweetPost) {
	tweetTime := tweet.TweetTime().Format("2006-01-02 15:04:05")

	t.Handler.Exec(
		"INSERT INTO tweet(tweet_id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		tweet.TweetID().ToString(), tweet.UserID().ToString(), tweet.Content(), tweetTime, tweetTime,
	)

	// Fanoutの呼び出しは非同期で実行するためレスポンスをまたない
	go event.Publish(tweet.UserID(), tweet.TweetID(), publisher.TweetPostPublisher{})
}
