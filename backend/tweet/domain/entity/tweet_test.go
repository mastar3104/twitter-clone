package entity

import (
	"github.com/google/uuid"
	"github.com/mastar3104/twitter-clone/tweet/domain/value"
	"testing"
)

func TestCreatTweet(t *testing.T) {
	t.Run("ユーザ情報と投稿内容から、ツイート情報が生成されること", func(t *testing.T) {
		user := User{
			userID:   value.UserIDReconstructor(uuid.New().String()),
			userName: "ツイートユーザ",
		}
		content := "ツイート内容"
		actual, event := CreatTweet(user, content)
		t.Run("ツイートイベントが発行されていること", func(t *testing.T) {
			if event == nil {
				t.Fatalf("Event is not occurring.")
			}
		})

		t.Run("tweetIDが発行されていること", func(t *testing.T) {
			if actual.tweetID.ToString() == "" {
				t.Fatalf("tweetID faild.")
			}
		})

		t.Run("引数で指定したユーザのuserIDが代入されていること", func(t *testing.T) {
			if actual.userID != user.UserID() {
				t.Fatalf("userID unmatch. actual: %s, expected: %s", actual.userID.ToString(), user.UserID().ToString())
			}
		})

		t.Run("引数で指定したcontentが代入されていること", func(t *testing.T) {
			if actual.content != content {
				t.Fatalf("content unmatch. actual: %s, expected: %s", actual.content, content)
			}
		})
	})
}
