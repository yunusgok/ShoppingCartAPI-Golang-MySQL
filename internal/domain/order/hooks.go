package order

import (
	"gorm.io/gorm"
	"picnshop/internal/domain/cart"
)

func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {

	var currentCart cart.Cart
	if err := tx.Where("UserID = ?", order.UserID).First(&currentCart).Error; err != nil {
		return err
	}
	if err := tx.Where("CartID = ?", currentCart.ID).Unscoped().Delete(&cart.Item{}).Error; err != nil {
		return err
	}

	if err := tx.Unscoped().Delete(&currentCart).Error; err != nil {
		return err
	}
	return nil
}
