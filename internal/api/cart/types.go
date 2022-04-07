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

type CreateProductRequest struct {
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float32 `json:"price"`
	Count      int     `json:"count"`
	CategoryID uint    `json:"categoryID"`
}
