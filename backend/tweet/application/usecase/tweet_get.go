package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/mastar3104/twitter-clone/tweet/application/query"
	"github.com/mastar3104/twitter-clone/tweet/application/view"
	"github.com/mastar3104/twitter-clone/tweet/domain/value"
)

type TweetGet struct {
	TweetQueryService query.TweetService
}

func (t TweetGet) Exec(context echo.Context) error {
	userID := value.UserIDReconstructor(context.Param("userId"))

	tweets := t.TweetQueryService.Get(userID)

	return context.JSON(200, view.CreateTweets(tweets))
}
