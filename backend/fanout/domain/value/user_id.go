package value

type UserID struct {
	value string
}

func (u UserID) ToString() string {
	return u.value
}

func UserIDReconstructor(value string) UserID {
	return UserID{value}
}
