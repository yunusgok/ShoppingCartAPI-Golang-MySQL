package cart

import (
	"gorm.io/gorm"
	"picnshop/internal/domain/product"
)

func (cartItem *Item) AfterSave(tx *gorm.DB) (err error) {

	var currentProduct product.Product
	if err := tx.Where("ID = ?", cartItem.ProductID).First(&currentProduct).Error; err != nil {
		return err
	}
	if err := tx.Model(&currentProduct).Update("StockCount", currentProduct.StockCount-cartItem.Count).Error; err != nil {
		return err
	}
	return
}
