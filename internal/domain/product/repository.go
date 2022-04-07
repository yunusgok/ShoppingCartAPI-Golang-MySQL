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
	return result.Error
}

func (r *Repository) FindBySKU(sku string) (*Product, error) {
	var product *Product
	err := r.db.Where("IsDeleted = ?", 0).Where(Product{SKU: sku}).First(&product).Error
	if err != nil {
		return nil, ErrProductNotFound
	}
	return product, nil
}

func (r *Repository) Create(p *Product) error {
	result := r.db.Create(p)

	return result.Error
}

func (r *Repository) GetAll(pageIndex, pageSize int) ([]Product, int) {
	var products []Product
	var count int64

	r.db.Where("IsDeleted = ?", 0).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)

	return products, int(count)
}

func (r *Repository) Delete(sku string) error {
	currentProduct, err := r.FindBySKU(sku)
	if err != nil {
		return err
	}
	currentProduct.IsDeleted = true

	err = r.Update(*currentProduct)
	return err
}
