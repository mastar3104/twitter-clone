package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mastar3104/twitter-clone/shared/infrastucture/adapter"
	"github.com/mastar3104/twitter-clone/user/domain/entity"
	"github.com/mastar3104/twitter-clone/user/domain/event"
	"github.com/mastar3104/twitter-clone/user/domain/value"
	"github.com/mastar3104/twitter-clone/user/infrastructure/publisher"
	"time"
)

type FollowRepository struct {
	Handler *adapter.MySQLHandler
}

func (f FollowRepository) Save(follow entity.Follow, event event.UserFollow) {
	currentTime := time.Now()
	datetimeStr := currentTime.Format("2006-01-02 15:04:05")

	f.Handler.Exec(
		"INSERT INTO follow(user_id, follow_user_id, created_at, updated_at) VALUES (?, ?, ?, ?)",
		follow.UserID().ToString(), follow.FollowUserID().ToString(), datetimeStr, datetimeStr,
	)

	// Fanoutの呼び出しは非同期で実行するためレスポンスをまたない
	go event.Publish(follow.UserID(), follow.FollowUserID(), publisher.UserFollowPublisher{})
}

func (f FollowRepository) Find(userID value.UserID, followUserID value.UserID) (entity.Follow, error) {
	rows := f.Handler.Query("SELECT user_id, follow_user_id FROM follow WHERE user_id = ? AND follow_user_id = ?", userID.ToString(), followUserID.ToString())
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err.Error())
		}
	}(rows)
	if rows.Next() {
		var userID string
		var followUserID string
		err := rows.Scan(&userID, &followUserID)
		if err != nil {
			panic(err.Error())
		}

		return entity.FollowReconstructor(
			value.UserIDReconstructor(userID),
			value.UserIDReconstructor(followUserID),
		), nil
	}
	return entity.Follow{}, errors.New(fmt.Sprintf("Follow with user_id is %s and follow_user_id is %s not found", userID.ToString(), followUserID.ToString()))
}
