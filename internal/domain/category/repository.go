package category

import (
	"gorm.io/gorm"
	"log"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) Migration() {
	err := r.db.AutoMigrate(&Category{})
	if err != nil {
		log.Print(err)
	}
	//https://gorm.io/docs/migration.html#content-inner
}

//TODO: create sample data from file
func (r *CategoryRepository) InsertSampleData() {
	cities := []Category{
		{Name: "P1", Desc: "product1"},
		{Name: "P2", Desc: "product1"},
	}

	for _, c := range cities {
		r.db.Where(Category{Name: c.Name}).Attrs(Category{Name: c.Name}).FirstOrCreate(&c)
	}
}

func (r *CategoryRepository) Create(c *Category) error {
	result := r.db.Create(c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CategoryRepository) GetByName(name string) []Category {
	var categories []Category
	r.db.Where("Name LIKE ?", "%"+name+"%").Find(&categories)

	return categories
}
