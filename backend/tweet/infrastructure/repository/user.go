package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mastar3104/twitter-clone/shared/infrastucture/adapter"
	"github.com/mastar3104/twitter-clone/tweet/domain/entity"
	"github.com/mastar3104/twitter-clone/tweet/domain/value"
)

type UserRepository struct {
	Handler *adapter.MySQLHandler
}

func (r UserRepository) Find(userID value.UserID) (entity.User, error) {
	rows := r.Handler.Query("SELECT user_name FROM user WHERE user_id = ?", userID.ToString())
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
