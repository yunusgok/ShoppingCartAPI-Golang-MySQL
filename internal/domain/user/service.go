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

//TODO: add validation middleware
func (c *Service) Create(user *User) error {
	existUser := c.r.GetByName(user.Username)
	if len(existUser) > 0 {
		return ErrUserExistWithName
	}

	err := c.r.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (c *Service) GetUser(username string, password string) (User, error) {
	users := c.r.GetByName(username)
	if len(users) == 0 {
		return User{}, ErrUserNotFound
	}
	user := users[0]
	match := hash.CheckPasswordHash(password+user.Salt, user.Password)
	if !match {
		return User{}, ErrUserNotFound
	}
	return user, nil
}
