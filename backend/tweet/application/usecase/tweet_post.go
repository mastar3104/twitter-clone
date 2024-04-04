package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/mastar3104/twitter-clone/tweet/application/view"
	"github.com/mastar3104/twitter-clone/tweet/domain/entity"
	"github.com/mastar3104/twitter-clone/tweet/domain/repository"
	"github.com/mastar3104/twitter-clone/tweet/domain/value"
)

type TweetPost struct {
	TweetRepository repository.TweetRepository
	UserRepository  repository.UserRepository
}

type TweetPostBody struct {
	Content string `json:"content"`
}

func (t TweetPost) Exec(context echo.Context) error {
	userID := value.UserIDReconstructor(context.Param("userId"))
	body := new(TweetPostBody)
	if err := context.Bind(body); err != nil || body.Content == "" {
		return context.JSON(400, view.Error{
			Msg: "content is required in the request body.",
		})
	}

	user, err := t.UserRepository.Find(userID)
	if err != nil {
		return context.JSON(404, view.Error{
			Msg: err.Error(),
		})
	}

	tweet, event := entity.CreatTweet(user, body.Content)
	t.TweetRepository.Save(tweet, *event)

	return context.JSON(200, view.Tweet{
		UserId:    user.UserID().ToString(),
		UserName:  user.UserName(),
		TweetId:   tweet.TweetID().ToString(),
		Content:   tweet.Content(),
		TweetTime: tweet.TweetTime().Unix(),
	})
}
