package order

import (
	"gorm.io/gorm"
	"picnshop/internal/domain/product"
	"picnshop/internal/domain/user"
)

type Order struct {
	gorm.Model
	UserID       uint
	User         user.User     `gorm:"foreignKey:ID;references:UserID"`
	OrderedItems []OrderedItem `gorm:"foreignKey:OrderID"`
	TotalPrice   float32
	IsCanceled   bool
}
type OrderedItem struct {
	gorm.Model
	Product   product.Product `gorm:"foreignKey:ProductID"`
	ProductID uint
	Count     int
	OrderID   uint
}

func NewOrder(uid uint, items []OrderedItem) *Order {
	var totalPrice float32 = 0.0
	for _, item := range items {
		totalPrice += item.Product.Price
	}
	return &Order{
		UserID:       uid,
		OrderedItems: items,
		TotalPrice:   totalPrice,
		IsCanceled:   false,
	}
}

func NewOrderedItem(count int, pid uint) *OrderedItem {
	return &OrderedItem{
		Count:     count,
		ProductID: pid,
	}
}
