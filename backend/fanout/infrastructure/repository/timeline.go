package repository

import (
	"github.com/mastar3104/twitter-clone/fanout/domain/entity"
	"github.com/mastar3104/twitter-clone/shared/infrastucture/adapter"
	"strings"
	"time"
)

type TimelineRepository struct {
	Handler *adapter.MySQLHandler
}

func (t TimelineRepository) Save(timelines []entity.Timeline) {
	if len(timelines) == 0 {
		return
	}
	currentTime := time.Now()
	datetimeStr := currentTime.Format("2006-01-02 15:04:05")

	query := "INSERT INTO timeline(user_id, tweet_id, created_at, updated_at) VALUES "
	prepared := make([]string, len(timelines))
	var values []interface{}
	for i, timeline := range timelines {
		prepared[i] = "(?, ?, ?, ?)"
		values = append(values, timeline.UserID().ToString(), timeline.TweetID().ToString(), datetimeStr, datetimeStr)
	}

	t.Handler.Exec(query+strings.Join(prepared, ","), values...)
}
