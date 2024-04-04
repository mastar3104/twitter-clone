package entity

import (
	"github.com/mastar3104/twitter-clone/user/domain/value"
	"testing"
)

func TestNewFollow(t *testing.T) {
	t.Run("引数に指定したuserがfollowUserをフォローしたentityを生成する", func(t *testing.T) {
		user := CreateUser("フォロワー")
		followUser := CreateUser("フォロー")
		actual, err := newFollow(user, followUser)

		t.Run("errorが発生していないこと", func(t *testing.T) {
			if err != nil {
				t.Fatalf(err.Error())
			}
		})

		t.Run("follow.UserID()とuser.UserID()が一致すること", func(t *testing.T) {
			if actual.UserID() != user.UserID() {
				t.Fatalf("follow.userID unmatch. actual:%s\n expected:%s", actual.UserID(), user.UserID())
			}
		})

		t.Run("follow.FollowUserID()とuser.UserID()が一致すること", func(t *testing.T) {
			if actual.FollowUserID() != followUser.UserID() {
				t.Fatalf("follow.followUserID unmatch. actual:%s\n expected:%s", actual.FollowUserID(), followUser.UserID())
			}
		})
	})

	t.Run("引数に指定したuserとfollowUserのIDが同じである場合、errorを返すこと", func(t *testing.T) {
		userID := value.CreateNewUserID()
		user := User{
			userID:   userID,
			userName: "フォロワー",
		}
		followUser := User{
			userID:   userID,
			userName: "フォロー",
		}
		actual, err := newFollow(user, followUser)

		t.Run("errorが発生すること", func(t *testing.T) {
			if err == nil {
				t.Fatalf("No errors have occurred.")
			}
		})

		t.Run("follow.UserID()がゼロ値であること", func(t *testing.T) {
			if actual.UserID().ToString() != "" {
				t.Fatalf("follow.userID is not Zero value. actual:%s", actual.UserID())
			}
		})

		t.Run("follow.FollowUserID()がゼロ値であること", func(t *testing.T) {
			if actual.FollowUserID().ToString() != "" {
				t.Fatalf("follow.followUserID is not Zero value. actual:%s", actual.UserID())
			}
		})
	})
}
