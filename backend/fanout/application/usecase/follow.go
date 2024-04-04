package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/mastar3104/twitter-clone/fanout/domain/entity"
	"github.com/mastar3104/twitter-clone/fanout/domain/repository"
	"github.com/mastar3104/twitter-clone/fanout/domain/value"
)

type Follow struct {
	TweetRepository    repository.TweetRepository
	TimelineRepository repository.TimelineRepository
}

func (f Follow) Exec(context echo.Context) error {
	userID := value.UserIDReconstructor(context.Param("userId"))
	followUserID := value.UserIDReconstructor(context.Param("followUserId"))

	tweets := f.TweetRepository.Get(followUserID)

	timelines := entity.CreateTimelineByFollow(tweets, userID)
	f.TimelineRepository.Save(timelines)

	return context.NoContent(200)
}
