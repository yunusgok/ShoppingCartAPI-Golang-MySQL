package user

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&User{})
	if err != nil {
		return
	}
	//https://gorm.io/docs/migration.html#content-inner
}

func (r *Repository) Create(u *User) error {
	result := r.db.Create(u)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//TODO it should return 1 result
func (r *Repository) GetByName(name string) []User {
	var users []User
	r.db.First(&users, "UserName = ?", name)
	return users
}

func (r *Repository) InsertSampleData() {
	user := NewUser("admin", "admin")

	err := r.db.Where(User{Username: user.Username}).Attrs(User{Username: user.Username, Password: user.Password}).FirstOrCreate(&user).Error
	if err != nil {
		return
	}
}
