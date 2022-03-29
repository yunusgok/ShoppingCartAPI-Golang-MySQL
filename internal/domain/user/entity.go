package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(100);unique_index"`
	Password  string
	Hash      string
	Salt      string
	Roles     []string
	Token     string `gorm:"unique"`
	IsDeleted bool
}

func NewUser(username string, password string, roles []string, token string) *User {

	return &User{
		Username:  username,
		Password:  password,
		Roles:     roles,
		Token:     token,
		IsDeleted: false,
	}
}
