package repository

import "github.com/mastar3104/twitter-clone/fanout/domain/entity"

type TimelineRepository interface {
	Save(timelines []entity.Timeline)
}
