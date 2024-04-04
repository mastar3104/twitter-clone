package entity

import (
	"github.com/mastar3104/twitter-clone/fanout/domain/value"
)

type Follower struct {
	userID         value.UserID
	followerUserID value.UserID
}

// UserID ユーザIDのGetter
func (u Follower) UserID() value.UserID {
	return u.userID
}

// FollowerUserID フォローユーザIDのGetter
func (u Follower) FollowerUserID() value.UserID {
	return u.followerUserID
}

// FollowReconstructor repository等外部の値からEntityを再生成するために利用
func FollowReconstructor(userID value.UserID, followerUserID value.UserID) Follower {
	return Follower{
		userID:         userID,
		followerUserID: followerUserID,
	}
}
