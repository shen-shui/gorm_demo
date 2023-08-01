package models

type User struct {
	Username string
	Id       int
	Age      int
}

func (User) TableName() string {
	return "user"
}
