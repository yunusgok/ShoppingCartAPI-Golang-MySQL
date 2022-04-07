package cart

import (
	"gorm.io/gorm"
	"picnshop/internal/domain/product"
	"picnshop/internal/domain/user"
)

type Cart struct {
	gorm.Model
	UserID uint
	User   user.User `gorm:"foreignKey:ID;references:UserID"`
}

func NewCart(uid uint) *Cart {
	return &Cart{
		UserID: uid,
	}
}

type Item struct {
	gorm.Model
	Product   product.Product `gorm:"foreignKey:ProductID"`
	ProductID uint
	Count     int
	CartID    uint
	Cart      Cart `gorm:"foreignKey:CartID" json:"-"`
}

func NewCartItem(productId uint, cartId uint, count int) *Item {
	return &Item{
		ProductID: productId,
		Count:     count,
		CartID:    cartId,
	}
}
