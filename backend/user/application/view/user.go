package view

type User struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
}

type Users struct {
	Users []User `json:"users"`
}
