package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UUID   uuid.UUID `gorm:"column:id; type:varchar(225); not null; unique;"`
	UserID uint      `gorm:"column:user_id; not null;"`
	// Relationship
	User User `gorm:"foreignKey:UserID; references:ID;"` // Many to One relationship
}

func (o *Order) TableName() string {
	return "order"
}
