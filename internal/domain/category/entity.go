package category

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Desc     string
	Code     string `gorm:"type:varchar(100);unique_index"`
	IsActive bool
}

func NewCategory(name string, desc string, code string) *Category {
	return &Category{
		Name:     name,
		Desc:     desc,
		Code:     code,
		IsActive: true,
	}
}
