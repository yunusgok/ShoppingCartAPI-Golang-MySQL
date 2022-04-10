package cart

import (
	"gorm.io/gorm"
	"picnshop/internal/domain/product"
)

func (cartItem *Item) BeforeSave(tx *gorm.DB) (err error) {

	var currentProduct product.Product
	var currentItem Item
	if err := tx.Where("ID = ?", cartItem.ProductID).First(&currentProduct).Error; err != nil {
		return err
	}
	oldCount := 0
	if err := tx.Where("ID = ?", cartItem.CartID).First(&currentItem).Error; err == nil {
		oldCount = currentItem.Count
	}
	newCount := currentProduct.StockCount + oldCount - cartItem.Count
	if newCount == 0 {
		if err := tx.Model(&currentProduct).Delete(currentItem).Error; err != nil {
			return err
		}
		return
	}
	if err := tx.Model(&currentProduct).Update("StockCount", newCount).Error; err != nil {
		return err
	}
	return
}
