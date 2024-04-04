package query

import (
	"database/sql"
	"github.com/mastar3104/twitter-clone/shared/infrastucture/adapter"
	"github.com/mastar3104/twitter-clone/tweet/application/model"
	"github.com/mastar3104/twitter-clone/tweet/domain/value"
	"time"
)

type TweetService struct {
	Handler *adapter.MySQLHandler
}

func (t TweetService) Get(userID value.UserID) []model.Tweet {
	rows := t.Handler.Query("SELECT t.user_id, u.user_name, t.tweet_id, t.content, t.created_at FROM tweet AS t INNER JOIN user AS u ON t.user_id = u.user_id WHERE t.user_id = ? ORDER BY t.created_at DESC", userID.ToString())
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err.Error())
		}
	}(rows)

	var tweets []model.Tweet
	for rows.Next() {
		var (
			userID    string
			userName  string
			tweetID   string
			content   string
			createdAt string
		)
		if err := rows.Scan(&userID, &userName, &tweetID, &content, &createdAt); err != nil {
			panic(err.Error())
		}
		tweetTile, _ := time.Parse("2006-01-02 15:04:05", createdAt)
		tweets = append(tweets, model.Tweet{
			UserId:    value.UserIDReconstructor(userID),
			UserName:  userName,
			TweetId:   value.TweetIDReconstructor(tweetID),
			Content:   content,
			TweetTime: tweetTile,
		})
	}
	return tweets
}
