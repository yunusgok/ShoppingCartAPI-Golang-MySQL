package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(30)"`
	Password string `gorm:"type:varchar(100)"`
	Salt     string `gorm:"type:varchar(100)"`
	Token    string `gorm:"type:varchar(100)"`
	//TODO add roles
	IsDeleted bool
}

func NewUser(username string, password string) *User {

	return &User{
		Username:  username,
		Password:  password,
		IsDeleted: false,
	}
}
