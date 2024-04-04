package entity

import (
	"errors"
	"github.com/mastar3104/twitter-clone/user/domain/value"
)

type Follow struct {
	userID       value.UserID
	followUserID value.UserID
}

// UserID ユーザIDのGetter
func (u Follow) UserID() value.UserID {
	return u.userID
}

// FollowUserID フォローユーザIDのGetter
func (u Follow) FollowUserID() value.UserID {
	return u.followUserID
}

// newFollow Followを生成するFactoryメソッド
func newFollow(user User, followUser User) (Follow, error) {
	if user.userID == followUser.userID {
		return Follow{}, errors.New("cannot follow yourself")
	}
	return Follow{
		userID:       user.userID,
		followUserID: followUser.userID,
	}, nil
}

func FollowReconstructor(userID value.UserID, followUserID value.UserID) Follow {
	return Follow{
		userID:       userID,
		followUserID: followUserID,
	}
}
