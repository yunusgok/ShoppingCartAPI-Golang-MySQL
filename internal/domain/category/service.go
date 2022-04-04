package category

type CategoryService struct {
	r CategoryRepository
}

func NewCategoryService(r CategoryRepository) *CategoryService {
	r.Migration()
	r.InsertSampleData()
	return &CategoryService{
		r: r,
	}
}

func (c *CategoryService) Create(category *Category) error {
	existCity := c.r.GetByName(category.Name)
	if len(existCity) > 0 {
		return ErrCategoryExistWithName
	}

	err := c.r.Create(category)
	if err != nil {
		return err
	}

	return nil
}
