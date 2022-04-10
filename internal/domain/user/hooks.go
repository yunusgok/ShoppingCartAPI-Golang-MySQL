package user

import (
	"picnshop/pkg/hash"

	"gorm.io/gorm"
)

func (u *User) BeforeSave(tx *gorm.DB) (err error) {

	if u.Salt == "" {
		//Create random string a salt to add to password
		salt := hash.CreateSalt()
		//create a hashed string from given password and created salt
		hashPassword, err := hash.HashPassword(u.Password + salt)
		if err != nil {
			return nil
		}
		u.Password = hashPassword
		u.Salt = salt
	}

	return
}
