package entity

import (
	"github.com/mastar3104/twitter-clone/user/domain/value"
	"testing"
)

func TestCreateUser(t *testing.T) {
	t.Run("新規のIDを発行したユーザエンティティを生成すること", func(t *testing.T) {
		actual := CreateUser("テストユーザ１")
		t.Run("ユーザIDが生成されていること(UUIDのため値の検証まで行わない)", func(t *testing.T) {
			if actual.userID.ToString() == "" {
				t.Fatalf("id faild.")
			}
		})
		t.Run("引数にしていたユーザ名と一致すること", func(t *testing.T) {
			if actual.userName != "テストユーザ１" {
				t.Fatalf("userName unmatch. actual: %s, expected: %s", actual.userID.ToString(), "テストユーザ１")
			}
		})
	})
}

func TestUser_Follow(t *testing.T) {
	t.Run("引数に指定したユーザとのフォロー情報を生成すること", func(t *testing.T) {
		userID := value.CreateNewUserID()
		followUserID := value.CreateNewUserID()
		target := UserReconstructor(userID, "フォロワー")
		follow := UserReconstructor(followUserID, "フォロー")
		actual, event, err := target.Follow(follow)

		t.Run("errorが発生しないこと", func(t *testing.T) {
			if err != nil {
				t.Fatalf(err.Error())
			}
		})

		t.Run("フォローを契機としてドメインイベントが生成すること", func(t *testing.T) {
			if event == nil {
				t.Fatalf("Event is not occurring.")
			}
		})

		t.Run("フォローを行なったユーザのユーザIDと一致すること", func(t *testing.T) {
			if actual.userID.ToString() != userID.ToString() {
				t.Fatalf("userID unmatch. actual: %s, expected: %s", actual.userID.ToString(), userID.ToString())
			}
		})
		t.Run("フォローされたユーザのユーザIDがfollowUserIDと一致すること", func(t *testing.T) {
			if actual.followUserID.ToString() != followUserID.ToString() {
				t.Fatalf("userID unmatch. actual: %s, expected: %s", actual.followUserID.ToString(), followUserID.ToString())
			}
		})
	})

	t.Run("ユーザ自身をフォローした場合、エラーは発生すること", func(t *testing.T) {
		userId := value.CreateNewUserID()
		target := UserReconstructor(userId, "フォロワー")
		actual, event, err := target.Follow(target)

		t.Run("自分自身をフォローできないエラーが発生すること", func(t *testing.T) {
			if err == nil {
				t.Fatalf("No errors have occurred.")
			}
		})

		t.Run("フォローを契機としてドメインイベントが生成されないこと", func(t *testing.T) {
			if event != nil {
				t.Fatalf("Event is occurring.")
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
