package application

/**
 * 将来的にはMQコンシューマ等、イベントトリガーに相応しい機能へ移行予定
 */

import (
	"github.com/labstack/echo/v4"
	"github.com/mastar3104/twitter-clone/fanout/application/usecase"
	"github.com/mastar3104/twitter-clone/fanout/infrastructure/repository"
	"github.com/mastar3104/twitter-clone/shared/infrastucture/adapter"
)

func Route(echoServer *echo.Echo, handler *adapter.MySQLHandler) {

	// DI
	timelineRepository := repository.TimelineRepository{Handler: handler}
	followerRepository := repository.FollowerRepository{Handler: handler}
	tweetRepository := repository.TweetRepository{Handler: handler}
	tweetPost := usecase.TweetPost{
		TimelineRepository: timelineRepository,
		FollowerRepository: followerRepository,
	}
	follow := usecase.Follow{
		TimelineRepository: timelineRepository,
		TweetRepository:    tweetRepository,
	}

	echoServer.POST("/v1/users/:userId/tweets/:tweetId", tweetPost.Exec)

	echoServer.POST("/v1/users/:userId/follow/:followUserId", follow.Exec)

}
