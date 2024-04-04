package value

type UserID struct {
	value string
}

func (u UserID) ToString() string {
	return u.value
}

// UserIDReconstructor repository等外部の値からVOを再生成するために利用
func UserIDReconstructor(value string) UserID {
	return UserID{value}
}
