package po

import (
	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	ID        int64  `gorm:"column:id; type:int; not null; primaryKey; autoIncrement; comment: 'Unique identifier for a brand';"`
	BrandName string `gorm:"column:brand_name; unique; not null; comment: 'Brand name'"`
	BrandNote string `gorm:"column:brand_note; type:text; comment: 'Brand note'"`

	// Relationship
	Products []Product `gorm:"foreignKey:ProductBrandID;references:ID"`
}

func (b *Brand) TableName() string {
	return "brand"
}
