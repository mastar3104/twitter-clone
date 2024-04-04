package entity

import (
	"github.com/mastar3104/twitter-clone/user/domain/event"
	"github.com/mastar3104/twitter-clone/user/domain/value"
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

// Follow ユーザが引数に指定したユーザをフォローする
func (u User) Follow(followUser User) (Follow, *event.UserFollow, error) {
	follow, err := newFollow(u, followUser)
	if err != nil {
		return Follow{}, nil, err
	}
	return follow, &event.UserFollow{}, nil
}

// CreateUser 新規ユーザの発行
func CreateUser(userName string) User {
	return User{
		userID:   value.CreateNewUserID(),
		userName: userName,
	}
}

// UserReconstructor repository等外部の値からEntityを再生成するために利用
func UserReconstructor(userID value.UserID, userName string) User {
	return User{
		userID:   userID,
		userName: userName,
	}
}
