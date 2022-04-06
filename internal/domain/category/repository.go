package category

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Category{})
	if err != nil {
		log.Print(err)
	}
	//https://gorm.io/docs/migration.html#content-inner
}

//TODO: create sample data from file
func (r *Repository) InsertSampleData() {
	categories := []Category{
		{Name: "P1", Desc: "product1"},
		{Name: "P2", Desc: "product1"},
	}

	for _, c := range categories {
		r.db.Where(Category{Name: c.Name}).Attrs(Category{Name: c.Name}).FirstOrCreate(&c)
	}
}

func (r *Repository) Create(c *Category) error {
	result := r.db.Create(c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) GetByName(name string) []Category {
	var categories []Category
	r.db.Where("Name LIKE ?", "%"+name+"%").Find(&categories)

	return categories
}

//TODO: bulk insert library exist look for implementation
func (r *Repository) BulkCreate(categories []*Category) {
	for _, c := range categories {
		r.db.Where(Category{Name: c.Name}).Attrs(Category{Name: c.Name}).FirstOrCreate(&c)
	}
}
