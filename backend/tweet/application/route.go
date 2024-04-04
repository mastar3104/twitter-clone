package application

import (
	"github.com/labstack/echo/v4"
	"github.com/mastar3104/twitter-clone/shared/infrastucture/adapter"
	"github.com/mastar3104/twitter-clone/tweet/application/usecase"
	"github.com/mastar3104/twitter-clone/tweet/infrastructure/query"
	"github.com/mastar3104/twitter-clone/tweet/infrastructure/repository"
)

func Route(echoServer *echo.Echo, handler *adapter.MySQLHandler) {

	/* DI */
	userRepository := repository.UserRepository{Handler: handler}
	tweetRepository := repository.TweetRepository{Handler: handler}
	tweetQueryService := query.TweetService{Handler: handler}
	timelineQueryService := query.TimelineService{Handler: handler}
	tweetGetUseCase := usecase.TweetGet{TweetQueryService: tweetQueryService}
	tweetPostUseCase := usecase.TweetPost{
		UserRepository:  userRepository,
		TweetRepository: tweetRepository,
	}
	timelineGet := usecase.TimelineGet{TimelineQueryService: timelineQueryService}

	echoServer.POST("/v1/tweets/:userId", tweetPostUseCase.Exec)

	echoServer.GET("/v1/tweets/:userId", tweetGetUseCase.Exec)

	echoServer.GET("/v1/timeline/:userId", timelineGet.Exec)
}
