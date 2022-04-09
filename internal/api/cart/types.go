package cart

type ItemCartRequest struct {
	UserId uint   `json:"userId"`
	SKU    string `json:"sku"`
	Count  int    `json:"count"`
}

type CreateCategoryResponse struct {
	Message string `json:"message"`
}
