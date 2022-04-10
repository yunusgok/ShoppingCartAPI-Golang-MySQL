package order

type CompleteOrderRequest struct {
	UserId uint `json:"userId"`
}
type CancelOrderRequest struct {
	UserId  uint `json:"userId"`
	OrderId uint `json:"orderId"`
}
