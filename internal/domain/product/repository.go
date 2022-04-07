package product

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Product{})
	if err != nil {
		log.Print(err)
	}
}

func (r *Repository) Update(product Product) error {
	result := r.db.Save(product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) FindBySKU(sku string) (*Product, error) {
	var product *Product
	err := r.db.Where(Product{SKU: sku}).First(&product).Error
	if err != nil {
		return nil, ErrProductNotFound
	}
	return product, nil
}

func (r *Repository) Create(p *Product) error {
	result := r.db.Create(p)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
