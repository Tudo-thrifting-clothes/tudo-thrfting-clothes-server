package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"column:id; type:varchar(225); not null; unique; primary_key;  index: idx_uuid;"`
	UserName string    `gorm:"column:user_name"`
	Email    string    `gorm:"column:email; type:varchar(225); not null; unique;"`
	IsActive bool      `gorm:"column:is_active; type:boolean; default:true;"`
	// Relationship
	Orders []Order `gorm:"foreignKey:UserID; references:ID;"`
}

func (u *User) TableName() string {
	return "user"
}
