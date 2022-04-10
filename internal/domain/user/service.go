package user

import "picnshop/pkg/hash"

type Service struct {
	r Repository
}

func NewUserService(r Repository) *Service {
	r.Migration()
	r.InsertSampleData()
	return &Service{
		r: r,
	}
}

func (c *Service) Create(user *User) error {
	if user.Password != user.Password2 {
		return ErrMismatchedPasswords
	}
	_, err := c.r.GetByName(user.Username)
	if err == nil {
		return ErrUserExistWithName
	}
	if ValidateUserName(user.Username) {
		return ErrInvalidUsername
	}
	if ValidatePassword(user.Password) {
		return ErrInvalidPassword
	}
	err = c.r.Create(user)
	return err
}

func (c *Service) GetUser(username string, password string) (User, error) {
	user, err := c.r.GetByName(username)
	if err != nil {
		return User{}, ErrUserNotFound
	}
	match := hash.CheckPasswordHash(password+user.Salt, user.Password)
	if !match {
		return User{}, ErrUserNotFound
	}
	return user, nil
}
func (c *Service) UpdateUser(user *User) error {
	return c.r.Update(user)
}
