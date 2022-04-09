package order

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Order{})
	if err != nil {
		log.Print(err)
	}
}

func (r *Repository) Update(newOrder Order) error {
	result := r.db.Save(newOrder)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) Create(ci *Order) error {
	result := r.db.Create(ci)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
