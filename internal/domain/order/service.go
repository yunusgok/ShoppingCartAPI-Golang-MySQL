package order

import (
	"picnshop/internal/domain/cart"
	"picnshop/internal/domain/product"
	"picnshop/pkg/pagination"
	"time"
)

var day14ToHours float64 = 336

type Service struct {
	orderRepository       Repository
	orderedItemRepository OrderedItemRepository
	productRepository     product.Repository
	cartRepository        cart.Repository
	cartItemRepository    cart.ItemRepository
}

func NewService(
	orderRepository Repository,
	orderedItemRepository OrderedItemRepository,
	productRepository product.Repository,
	cartRepository cart.Repository,
	cartItemRepository cart.ItemRepository,
) *Service {
	orderRepository.Migration()
	orderedItemRepository.Migration()
	return &Service{
		orderRepository:       orderRepository,
		orderedItemRepository: orderedItemRepository,
		productRepository:     productRepository,
		cartRepository:        cartRepository,
		cartItemRepository:    cartItemRepository,
	}

}

// CompleteOrder creates an order with user's cart if user has a cart and there are items in it
func (c *Service) CompleteOrder(userId uint) error {
	currentCart, err := c.cartRepository.FindOrCreateByUserID(userId)
	if err != nil {
		return err
	}
	cartItems, err := c.cartItemRepository.GetItems(currentCart.UserID)
	if err != nil {
		return err
	}
	if len(cartItems) == 0 {
		return ErrEmptyCartFound
	}
	orderedItems := make([]OrderedItem, 0)
	for _, item := range cartItems {
		orderedItems = append(orderedItems, *NewOrderedItem(item.Count, item.ProductID))
	}
	err = c.orderRepository.Create(NewOrder(userId, orderedItems))
	return err
}

// CancelOrder cancels the order with oid(OrderID) if order is exist and  14 day is not passed
// Updates Order's IsCanceled field as true
func (c *Service) CancelOrder(uid, oid uint) error {
	currentOrder, err := c.orderRepository.FindByOrderID(oid)
	if err != nil {
		return err
	}
	if currentOrder.UserID != uid {
		return ErrInvalidOrderID
	}
	if currentOrder.CreatedAt.Sub(time.Now()).Hours() > day14ToHours {
		return ErrCancelDurationPassed
	}
	currentOrder.IsCanceled = true
	err = c.orderRepository.Update(*currentOrder)

	return err
}

// GetAll return all Orders of the user in given page
func (c *Service) GetAll(page *pagination.Pages, uid uint) *pagination.Pages {
	orders, count := c.orderRepository.GetAll(page.Page, page.PageSize, uid)
	page.Items = orders
	page.TotalCount = count
	return page

}
