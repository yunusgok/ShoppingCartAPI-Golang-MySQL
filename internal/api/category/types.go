package category

type CreateCategoryRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type CategoryResponse struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
