package category

import "gorm.io/gorm"

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) Migration() {
	r.db.AutoMigrate(&Category{})
	//https://gorm.io/docs/migration.html#content-inner
}

func (r *CategoryRepository) InsertSampleData() {
	cities := []Category{
		{Name: "P1", Desc: "product1"},
		{Name: "P2", Desc: "product1"},
	}

	for _, c := range cities {
		r.db.Where(Category{Name: c.Name}).Attrs(Category{Name: c.Name}).FirstOrCreate(&c)
	}
}