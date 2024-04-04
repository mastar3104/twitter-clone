package entity

import (
	"github.com/mastar3104/twitter-clone/tweet/domain/value"
)

type User struct {
	userID   value.UserID
	userName string
}

// UserID ユーザIDのGetter
func (u User) UserID() value.UserID {
	return u.userID
}

// UserName ユーザ名のGetter
func (u User) UserName() string {
	return u.userName
}

// UserReconstructor repository等外部の値からEntityを再生成するために利用
func UserReconstructor(userID value.UserID, userName string) User {
	return User{
		userID:   userID,
		userName: userName,
	}
}
