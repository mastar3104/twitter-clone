package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mastar3104/twitter-clone/shared/infrastucture/adapter"
	"github.com/mastar3104/twitter-clone/user/domain/entity"
	"github.com/mastar3104/twitter-clone/user/domain/value"
	"time"
)

type UserRepository struct {
	Handler *adapter.MySQLHandler
}

func (u UserRepository) Save(user entity.User) {
	currentTime := time.Now()
	datetimeStr := currentTime.Format("2006-01-02 15:04:05")

	u.Handler.Exec(
		"INSERT INTO user(user_id, user_name, created_at, updated_at) VALUES (?, ?, ?, ?)",
		user.UserID().ToString(), user.UserName(), datetimeStr, datetimeStr,
	)
}

func (u UserRepository) Find(userID value.UserID) (entity.User, error) {
	rows := u.Handler.Query("SELECT user_name FROM user WHERE user_id = ?", userID.ToString())
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err.Error())
		}
	}(rows)
	if rows.Next() {
		var userName string
		err := rows.Scan(&userName)
		if err != nil {
			panic(err.Error())
		}

		return entity.UserReconstructor(userID, userName), nil
	}
	return entity.User{}, errors.New(fmt.Sprintf("User with user_id %s not found", userID.ToString()))
}

func (u UserRepository) Get() []entity.User {
	rows := u.Handler.Query("SELECT user_id, user_name FROM user")
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err.Error())
		}
	}(rows)
	var users []entity.User
	for rows.Next() {
		var userID string
		var userName string
		err := rows.Scan(&userID, &userName)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, entity.UserReconstructor(value.UserIDReconstructor(userID), userName))
	}
	return users
}
