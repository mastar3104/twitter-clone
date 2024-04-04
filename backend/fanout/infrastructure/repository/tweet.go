package repository

import (
	"database/sql"
	"github.com/mastar3104/twitter-clone/fanout/domain/entity"
	"github.com/mastar3104/twitter-clone/fanout/domain/value"
	"github.com/mastar3104/twitter-clone/shared/infrastucture/adapter"
	"time"
)

type TweetRepository struct {
	Handler *adapter.MySQLHandler
}

func (t TweetRepository) Get(userID value.UserID) []entity.Tweet {
	rows := t.Handler.Query("SELECT tweet_id, content, created_at FROM tweet WHERE user_id = ?", userID.ToString())
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err.Error())
		}
	}(rows)
	var tweets []entity.Tweet
	for rows.Next() {
		var (
			tweetID   string
			content   string
			createdAt string
		)
		if err := rows.Scan(&tweetID, &content, &createdAt); err != nil {
			panic(err.Error())
		}
		tweetTile, _ := time.Parse("2006-01-02 15:04:05", createdAt)
		tweets = append(tweets, entity.TweetReconstructor(
			value.TweetIDReconstructor(tweetID),
			userID,
			content,
			tweetTile,
		))
	}
	return tweets
}
