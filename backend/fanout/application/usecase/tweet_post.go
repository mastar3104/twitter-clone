package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/mastar3104/twitter-clone/fanout/domain/entity"
	"github.com/mastar3104/twitter-clone/fanout/domain/repository"
	"github.com/mastar3104/twitter-clone/fanout/domain/value"
)

type TweetPost struct {
	FollowerRepository repository.FollowerRepository
	TimelineRepository repository.TimelineRepository
}

func (t TweetPost) Exec(context echo.Context) error {
	userID := value.UserIDReconstructor(context.Param("userId"))
	tweetID := value.TweetIDReconstructor(context.Param("tweetId"))

	followers := t.FollowerRepository.Get(userID)

	timelines := entity.CreateTimelineByTweetPost(userID, followers, tweetID)
	t.TimelineRepository.Save(timelines)

	return context.NoContent(200)
}
