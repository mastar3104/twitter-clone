package value

type TweetID struct {
	value string
}

func (t TweetID) ToString() string {
	return t.value
}

// TweetIDReconstructor repository等外部の値からVOを再生成するために利用
func TweetIDReconstructor(value string) TweetID {
	return TweetID{value}
}
