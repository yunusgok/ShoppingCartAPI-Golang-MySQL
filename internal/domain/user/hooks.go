package user

import (
	"picnshop/pkg/hash"

	"gorm.io/gorm"
)

func (u *User) BeforeSave(tx *gorm.DB) (err error) {

	if u.Password != "" {
		//Create random string a salt to add to password
		salt := hash.CreateSalt()
		//create a hashed string from given password and created salt
		hash, err := hash.HashPassword(u.Password + salt)
		if err != nil {
			return nil
		}
		u.Password = hash
		u.Salt = salt
	}

	return
}
