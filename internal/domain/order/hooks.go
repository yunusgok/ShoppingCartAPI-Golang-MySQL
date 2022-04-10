package order

import (
	"gorm.io/gorm"
	"picnshop/internal/domain/cart"
	"picnshop/internal/domain/product"
)

// BeforeCreate of Order finds cart of the user and deletes inside the cart.
// After deletion of items,it deletes cart
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

// BeforeSave of OrderedItem updates products' stock count
func (orderedItem *OrderedItem) BeforeSave(tx *gorm.DB) (err error) {

	var currentProduct product.Product
	var currentOrderedItem OrderedItem
	if err := tx.Where("ID = ?", orderedItem.ProductID).First(&currentProduct).Error; err != nil {
		return err
	}
	reservedStockCount := 0
	if err := tx.Where("ID = ?", orderedItem.ID).First(&currentOrderedItem).Error; err == nil {
		reservedStockCount = currentOrderedItem.Count
	}
	newStockCount := currentProduct.StockCount + reservedStockCount - orderedItem.Count
	if newStockCount < 0 {
		return ErrNotEnoughStock
	}
	if err := tx.Model(&currentProduct).Update("StockCount", newStockCount).Error; err != nil {
		return err
	}
	if orderedItem.Count == 0 {
		err := tx.Unscoped().Delete(currentOrderedItem).Error
		return err
	}
	return
}

// BeforeUpdate of Order checks if an order is canceled.
// If Order is canceled the amount will be returned to products' stockCount
func (order *Order) BeforeUpdate(tx *gorm.DB) (err error) {

	if order.IsCanceled {
		var orderedItems []OrderedItem
		if err := tx.Where("OrderID = ?", order.ID).Find(&orderedItems).Error; err != nil {
			return err
		}
		for _, item := range orderedItems {
			var currentProduct product.Product
			if err := tx.Where("ID = ?", item.ProductID).First(&currentProduct).Error; err != nil {
				return err
			}
			newStockCount := currentProduct.StockCount + item.Count
			if err := tx.Model(&currentProduct).Update(
				"StockCount", newStockCount).Error; err != nil {
				return err
			}
			if err := tx.Model(&item).Update(
				"IsCanceled", true).Error; err != nil {
				return err
			}
		}
	}
	return

}
