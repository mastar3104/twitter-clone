package entity

import "github.com/mastar3104/twitter-clone/fanout/domain/value"

type Timeline struct {
	userID  value.UserID
	tweetID value.TweetID
}

func (t Timeline) UserID() value.UserID {
	return t.userID
}

func (t Timeline) TweetID() value.TweetID {
	return t.tweetID
}

// CreateTimelineByTweetPost ツイート投稿を契機としてタイムライン情報の生成
func CreateTimelineByTweetPost(userID value.UserID, followers []Follower, tweetID value.TweetID) []Timeline {
	timelines := make([]Timeline, len(followers)+1)
	timelines[0] = Timeline{
		userID:  userID,
		tweetID: tweetID,
	}
	for i, follow := range followers {
		timelines[i+1] = Timeline{
			userID:  follow.followerUserID,
			tweetID: tweetID,
		}
	}
	return timelines
}

// CreateTimelineByFollow ユーザの新規フォローを契機としてタイムライン情報の生成
func CreateTimelineByFollow(tweets []Tweet, userID value.UserID) []Timeline {
	timelines := make([]Timeline, len(tweets))
	for i, tweet := range tweets {
		timelines[i] = Timeline{
			userID:  userID,
			tweetID: tweet.tweetID,
		}
	}
	return timelines
}
