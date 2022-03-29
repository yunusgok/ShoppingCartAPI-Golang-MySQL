package product

import (
	"picnshop/internal/domain/category"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name       string
	SKU        string `gorm:"type:varchar(100);unique_index"`
	Desc       string
	StockCount int
	Price      float32
	CategoryID uint
	Category   category.Category `gorm:"foreignKey:CategoryID"`
}

func NewProduct(name string, sku string, desc string, stockCount int, price float32, categoryId uint) *Product {
	return &Product{
		Name:       name,
		SKU:        sku,
		Desc:       desc,
		StockCount: stockCount,
		Price:      price,
		CategoryID: categoryId,
	}
}