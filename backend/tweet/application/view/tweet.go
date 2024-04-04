package view

import "github.com/mastar3104/twitter-clone/tweet/application/model"

type Tweet struct {
	UserId    string `json:"userId"`
	UserName  string `json:"userName"`
	TweetId   string `json:"tweetId"`
	Content   string `json:"content"`
	TweetTime int64  `json:"tweetTime"`
}

type Tweets struct {
	Tweets []Tweet `json:"tweets"`
}

func CreateTweets(tweets []model.Tweet) Tweets {
	tweetsView := make([]Tweet, len(tweets))

	for i, tweet := range tweets {
		tweetsView[i] = Tweet{
			UserId:    tweet.UserId.ToString(),
			UserName:  tweet.UserName,
			TweetId:   tweet.TweetId.ToString(),
			Content:   tweet.Content,
			TweetTime: tweet.TweetTime.Unix(),
		}
	}

	return Tweets{
		Tweets: tweetsView,
	}
}
