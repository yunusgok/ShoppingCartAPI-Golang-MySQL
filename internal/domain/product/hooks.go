package product

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BeforeSave of Product a unique sku number will be created and saved to db
func (p *Product) BeforeSave(tx *gorm.DB) (err error) {

	p.SKU = uuid.New().String()
	return
}
