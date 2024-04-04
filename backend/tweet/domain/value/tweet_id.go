package value

import "github.com/google/uuid"

type TweetID struct {
	value string
}

func (t TweetID) ToString() string {
	return t.value
}

// CreateTweetID 新たにツイートの投稿を行う際のID発行
func CreateTweetID() TweetID {
	return TweetID{
		value: uuid.New().String(),
	}
}

// TweetIDReconstructor repository等外部の値からVOを再生成するために利用
func TweetIDReconstructor(value string) TweetID {
	return TweetID{value}
}
