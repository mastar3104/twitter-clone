package value

import "github.com/google/uuid"

type UserID struct {
	value string
}

func (u UserID) ToString() string {
	return u.value
}

func CreateNewUserID() UserID {
	return UserID{
		value: uuid.New().String(),
	}
}

func UserIDReconstructor(value string) UserID {
	return UserID{value}
}
