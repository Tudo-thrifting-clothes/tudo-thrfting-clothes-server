package po

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID           int64  `gorm:"column:id; type:int; not null; primaryKey; autoIncrement; comment: 'Unique identifier for a category';"`
	CategoryName string `gorm:"column:category_name; unique; not null; comment: 'category name'"`
	CategoryNote string `gorm:"column:category_note; type:text; comment: 'category note'"`

	// Relationships
	Products []Product `gorm:"foreignKey:ProductCategoryID;references:ID"`
}

func (c *Category) TableName() string {
	return "category"
}
