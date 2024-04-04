package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/mastar3104/twitter-clone/user/application/view"
	"github.com/mastar3104/twitter-clone/user/domain/repository"
)

type UsersGet struct {
	UserRepository repository.UserRepository
}

func (u UsersGet) Exec(context echo.Context) error {
	users := u.UserRepository.Get()

	var views []view.User
	for _, user := range users {
		views = append(views, view.User{
			UserID:   user.UserID().ToString(),
			UserName: user.UserName(),
		})
	}

	return context.JSON(200, view.Users{
		Users: views,
	})
}
