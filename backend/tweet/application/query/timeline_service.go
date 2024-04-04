package query

import (
	"github.com/mastar3104/twitter-clone/tweet/application/model"
	"github.com/mastar3104/twitter-clone/tweet/domain/value"
)

type TimelineService interface {
	Get(userID value.UserID) []model.Tweet
}
