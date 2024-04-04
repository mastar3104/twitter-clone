package entity

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mastar3104/twitter-clone/fanout/domain/value"
	"testing"
	"time"
)

func TestCreateTimelineByTweetPost(t *testing.T) {
	t.Run("ツイート投稿を契機に、該当ツイートのタイムラインをフォロワーと自分の生成する", func(t *testing.T) {
		userID := value.UserIDReconstructor(uuid.New().String())
		tweetID := value.TweetIDReconstructor(uuid.New().String())
		follower1ID := value.UserIDReconstructor(uuid.New().String())
		follower2ID := value.UserIDReconstructor(uuid.New().String())
		follower := []Follower{
			{
				userID:         userID,
				followerUserID: follower1ID,
			},
			{
				userID:         userID,
				followerUserID: follower2ID,
			},
		}
		actual := CreateTimelineByTweetPost(userID, follower, tweetID)
		expected := []Timeline{
			{
				userID:  userID,
				tweetID: tweetID,
			},
			{
				userID:  follower1ID,
				tweetID: tweetID,
			},
			{
				userID:  follower2ID,
				tweetID: tweetID,
			},
		}

		t.Run("生成されたタイムライン数が想定通りであること", func(t *testing.T) {
			if len(actual) != len(expected) {
				t.Fatalf("timeline size unmatched. actual:%d, expected:%d", len(actual), len(expected))
			}
		})
		for i, timeline := range actual {
			t.Run(fmt.Sprintf("[%d].userID が一致すること", i), func(t *testing.T) {
				if timeline.userID != expected[i].userID {
					t.Fatalf("[%d].userID unmatched. actual:%s, expected:%s", i, timeline.userID.ToString(), expected[i].userID.ToString())
				}
			})
			t.Run(fmt.Sprintf("[%d].tweetID が一致すること", i), func(t *testing.T) {
				if timeline.tweetID != expected[i].tweetID {
					t.Fatalf("[%d].tweetID unmatched. actual:%s, expected:%s", i, timeline.tweetID.ToString(), expected[i].tweetID.ToString())
				}
			})
		}
	})

	t.Run("一人もフォロワーがいない場合、自分のタイムラインのみ生成すること", func(t *testing.T) {
		userID := value.UserIDReconstructor(uuid.New().String())
		tweetID := value.TweetIDReconstructor(uuid.New().String())
		var followers []Follower
		actual := CreateTimelineByTweetPost(userID, followers, tweetID)
		expected := []Timeline{
			{
				userID:  userID,
				tweetID: tweetID,
			},
		}
		t.Run("生成されたタイムライン数が想定通りであること", func(t *testing.T) {
			if len(actual) != len(expected) {
				t.Fatalf("timeline size unmatched. actual:%d, expected:%d", len(actual), len(expected))
			}
		})
		for i, timeline := range actual {
			t.Run(fmt.Sprintf("[%d].userID が一致すること", i), func(t *testing.T) {
				if timeline.userID != expected[i].userID {
					t.Fatalf("[%d].userID unmatched. actual:%s, expected:%s", i, timeline.userID.ToString(), expected[i].userID.ToString())
				}
			})
			t.Run(fmt.Sprintf("[%d].tweetID が一致すること", i), func(t *testing.T) {
				if timeline.tweetID != expected[i].tweetID {
					t.Fatalf("[%d].tweetID unmatched. actual:%s, expected:%s", i, timeline.tweetID.ToString(), expected[i].tweetID.ToString())
				}
			})
		}
	})
}

func TestCreateTimelineByFollow(t *testing.T) {
	t.Run("フォローを契機に、フォローしたユーザのツイートを自分のタイムラインに生成する", func(t *testing.T) {
		userID := value.UserIDReconstructor(uuid.New().String())
		followUserID := value.UserIDReconstructor(uuid.New().String())
		followUserTweet1ID := value.TweetIDReconstructor(uuid.New().String())
		followUserTweet2ID := value.TweetIDReconstructor(uuid.New().String())
		tweets := []Tweet{
			{
				tweetID:   followUserTweet1ID,
				userID:    followUserID,
				content:   "投稿１",
				tweetTime: time.Now(),
			},
			{
				tweetID:   followUserTweet2ID,
				userID:    followUserID,
				content:   "投稿２",
				tweetTime: time.Now(),
			},
		}
		actual := CreateTimelineByFollow(tweets, userID)
		expected := []Timeline{
			{
				userID:  userID,
				tweetID: followUserTweet1ID,
			},
			{
				userID:  userID,
				tweetID: followUserTweet2ID,
			},
		}

		t.Run("生成されたタイムライン数が想定通りであること", func(t *testing.T) {
			if len(actual) != len(expected) {
				t.Fatalf("timeline size unmatched. actual:%d, expected:%d", len(actual), len(expected))
			}
		})
		for i, timeline := range actual {
			t.Run(fmt.Sprintf("[%d].userID が一致すること", i), func(t *testing.T) {
				if timeline.userID != expected[i].userID {
					t.Fatalf("[%d].userID unmatched. actual:%s, expected:%s", i, timeline.userID.ToString(), expected[i].userID.ToString())
				}
			})
			t.Run(fmt.Sprintf("[%d].tweetID が一致すること", i), func(t *testing.T) {
				if timeline.tweetID != expected[i].tweetID {
					t.Fatalf("[%d].tweetID unmatched. actual:%s, expected:%s", i, timeline.tweetID.ToString(), expected[i].tweetID.ToString())
				}
			})
		}

	})
	t.Run("フォローしたユーザに1件もツイートがない場合、タイムラインを生成しないこと", func(t *testing.T) {
		userID := value.UserIDReconstructor(uuid.New().String())
		var tweets []Tweet

		actual := CreateTimelineByFollow(tweets, userID)

		t.Run("タイムラインの件数が0件であること", func(t *testing.T) {
			if len(actual) != 0 {
				t.Fatalf("timeline size unmatched. actual:%d, expected:%d", len(actual), 0)
			}
		})
	})
}
