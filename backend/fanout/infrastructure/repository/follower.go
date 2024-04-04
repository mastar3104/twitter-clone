package repository

import (
	"database/sql"
	"github.com/mastar3104/twitter-clone/fanout/domain/entity"
	"github.com/mastar3104/twitter-clone/fanout/domain/value"
	"github.com/mastar3104/twitter-clone/shared/infrastucture/adapter"
)

type FollowerRepository struct {
	Handler *adapter.MySQLHandler
}

func (f FollowerRepository) Get(userID value.UserID) []entity.Follower {
	rows := f.Handler.Query("SELECT user_id FROM follow WHERE follow_user_id = ?", userID.ToString())
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err.Error())
		}
	}(rows)
	var followers []entity.Follower
	for rows.Next() {
		var (
			followerUserID string
		)
		if err := rows.Scan(&followerUserID); err != nil {
			panic(err.Error())
		}
		followers = append(followers, entity.FollowReconstructor(
			userID,
			value.UserIDReconstructor(followerUserID),
		))
	}
	return followers
}
