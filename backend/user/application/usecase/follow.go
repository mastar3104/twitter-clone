package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/mastar3104/twitter-clone/user/application/view"
	"github.com/mastar3104/twitter-clone/user/domain/repository"
	"github.com/mastar3104/twitter-clone/user/domain/value"
)

type Follow struct {
	FollowRepository repository.FollowRepository
	UserRepository   repository.UserRepository
}

type FollowPostBody struct {
	FollowUserID string `json:"followUserId"`
}

func (f Follow) Exec(context echo.Context) error {
	userID := context.Param("userId")

	body := new(FollowPostBody)
	if err := context.Bind(body); err != nil || body.FollowUserID == "" {
		return context.JSON(400, view.Error{
			Msg: "followUserID is required in the request body.",
		})
	}

	myID := value.UserIDReconstructor(userID)
	followUserID := value.UserIDReconstructor(body.FollowUserID)

	user, err := f.UserRepository.Find(myID)
	if err != nil {
		return context.JSON(404, view.Error{
			Msg: err.Error(),
		})
	}
	followUser, err := f.UserRepository.Find(followUserID)
	if err != nil {
		return context.JSON(404, view.Error{
			Msg: err.Error(),
		})
	}

	if _, err := f.FollowRepository.Find(myID, followUserID); err == nil {
		return context.JSON(400, view.Error{
			Msg: "already follow. followUserId: " + followUserID.ToString(),
		})
	}

	follow, event, err := user.Follow(followUser)
	if err != nil {
		return context.JSON(400, view.Error{
			Msg: err.Error(),
		})
	}

	f.FollowRepository.Save(follow, *event)

	return context.NoContent(200)
}
