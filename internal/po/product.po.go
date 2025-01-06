package po

import (
	"tudo-thrfting-clothes-server/pkg/util"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID                uuid.UUID `gorm:"column:id; type:varchar(225); not null; unique; primary_key; index: idx_uuid;"`
	Name              string    `gorm:"column:name; type:varchar(225); not null; unique;"`
	Price             int32     `gorm:"column:price; not null;"`
	Stock             int32     `gorm:"column:stock; not null;"`
	Slug              string    `gorm:"column:slug; type:varchar(225);"`
	Description       string    `gorm:"column:description; type:text;"`
	ProductBrandID    *int64    `gorm:"column:product_brand_id;"`    // Foreign key for Brand
	ProductCategoryID *int64    `gorm:"column:product_category_id;"` // Foreign key for Category

	// Relationships
	Brand    *Brand    `gorm:"foreignKey:ProductBrandID;references:ID"`
	Category *Category `gorm:"foreignKey:ProductCategoryID;references:ID"`
}

func (p *Product) TableName() string {
	return "product"
}

// BeforeCreate is a GORM hook that generates a UUID for the ID field before creating a new record.
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}

	p.Slug = util.GenerateSlug(p.Name)

	return nil
}
