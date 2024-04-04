package query

import (
	"database/sql"
	"github.com/mastar3104/twitter-clone/shared/infrastucture/adapter"
	"github.com/mastar3104/twitter-clone/tweet/application/model"
	"github.com/mastar3104/twitter-clone/tweet/domain/value"
	"time"
)

type TimelineService struct {
	Handler *adapter.MySQLHandler
}

func (t TimelineService) Get(userID value.UserID) []model.Tweet {
	rows := t.Handler.Query("SELECT t2.user_id, u.user_name, t2.tweet_id, t2.content, t2.created_at FROM timeline AS t1 INNER JOIN tweet AS t2 ON t1.tweet_id = t2.tweet_id INNER JOIN user AS u ON t2.user_id = u.user_id WHERE t1.user_id = ? ORDER BY t2.created_at DESC", userID.ToString())
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
