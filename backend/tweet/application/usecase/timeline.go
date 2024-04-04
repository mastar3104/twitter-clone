package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/mastar3104/twitter-clone/tweet/application/query"
	"github.com/mastar3104/twitter-clone/tweet/application/view"
	"github.com/mastar3104/twitter-clone/tweet/domain/value"
)

type TimelineGet struct {
	TimelineQueryService query.TimelineService
}

func (t TimelineGet) Exec(context echo.Context) error {
	userID := value.UserIDReconstructor(context.Param("userId"))

	tweets := t.TimelineQueryService.Get(userID)

	return context.JSON(200, view.CreateTweets(tweets))
}
