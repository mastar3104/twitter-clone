package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/mastar3104/twitter-clone/user/application/view"
	"github.com/mastar3104/twitter-clone/user/domain/entity"
	"github.com/mastar3104/twitter-clone/user/domain/repository"
)

type UserPost struct {
	UserRepository repository.UserRepository
}

type UserPostBody struct {
	UserName string `json:"userName"`
}

func (u UserPost) Exec(context echo.Context) error {
	body := new(UserPostBody)
	if err := context.Bind(body); err != nil || body.UserName == "" {
		return context.JSON(400, view.Error{
			Msg: "userName is required in the request body.",
		})
	}

	user := entity.CreateUser(body.UserName)
	u.UserRepository.Save(user)

	return context.JSON(200, view.User{
		UserID:   user.UserID().ToString(),
		UserName: user.UserName(),
	})
}
