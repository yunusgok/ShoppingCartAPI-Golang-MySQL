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

func (c *CategoryService) Create(city *Category) error {
	existCity := c.r.GetByName(city.Name)
	if len(existCity) > 0 {
		return ErrCategoryExistWithName
	}

	err := c.r.Create(city)
	if err != nil {
		return err
	}

	return nil
}
