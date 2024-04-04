package application

import (
	"github.com/labstack/echo/v4"
	"github.com/mastar3104/twitter-clone/shared/infrastucture/adapter"
	"github.com/mastar3104/twitter-clone/user/application/usecase"
	"github.com/mastar3104/twitter-clone/user/infrastructure/repository"
)

func Route(echoServer *echo.Echo, handler *adapter.MySQLHandler) {

	/* DI */
	userRepository := repository.UserRepository{Handler: handler}
	followRepository := repository.FollowRepository{Handler: handler}
	userGetUseCase := usecase.UserGet{UserRepository: userRepository}
	usersGetUseCase := usecase.UsersGet{UserRepository: userRepository}
	userPostUseCase := usecase.UserPost{UserRepository: userRepository}
	followUseCase := usecase.Follow{
		FollowRepository: followRepository,
		UserRepository:   userRepository,
	}

	echoServer.GET("/v1/users/:userId", userGetUseCase.Exec)

	echoServer.GET("/v1/users", usersGetUseCase.Exec)

	echoServer.POST("/v1/users", userPostUseCase.Exec)

	echoServer.POST("/v1/users/:userId/follow", followUseCase.Exec)
}
