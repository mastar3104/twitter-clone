package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/mastar3104/twitter-clone/user/application/view"
	"github.com/mastar3104/twitter-clone/user/domain/repository"
	"github.com/mastar3104/twitter-clone/user/domain/value"
)

type UserGet struct {
	UserRepository repository.UserRepository
}

func (u UserGet) Exec(context echo.Context) error {
	userID := context.Param("userId")

	user, err := u.UserRepository.Find(value.UserIDReconstructor(userID))
	if err != nil {
		return context.JSON(404, view.Error{
			Msg: err.Error(),
		})
	}

	return context.JSON(200, view.User{
		UserID:   user.UserID().ToString(),
		UserName: user.UserName(),
	})
}
