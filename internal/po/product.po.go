package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID                uuid.UUID `gorm:"column:id; type:varchar(225); not null; unique; primary_key; index: idx_uuid;"`
	ProductName       string    `gorm:"column:product_name; type:varchar(225); not null;"`
	Price             int32     `gorm:"column:price; not null;"`
	Stock             int32     `gorm:"column:stock; not null;"`
	Description       string    `gorm:"column:description; type:text;"`
	ProductBrandID    int64     `gorm:"column:product_brand_id;"`    // Foreign key for Brand
	ProductCategoryID int64     `gorm:"column:product_category_id;"` // Foreign key for Category

	// Relationships
	Brand    Brand    `gorm:"foreignKey:ProductBrandID;references:ID"`
	Category Category `gorm:"foreignKey:ProductCategoryID;references:ID"`
}

func (p *Product) TableName() string {
	return "product"
}
