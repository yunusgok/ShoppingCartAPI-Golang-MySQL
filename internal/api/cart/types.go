package cart

type AddItemToCartRequest struct {
	UserId uint   `json:"userId"`
	SKU    string `json:"sku"`
	Count  int    `json:"count"`
}

type CreateCategoryResponse struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
