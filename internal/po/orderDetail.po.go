package po

// import (
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	gorm.Model
// 	UUID     uuid.UUID `gorm:"column:uuid; type:varchar(225); not null; unique;  index: idx_uuid;"`
// 	UserName string    `gorm:"column:user_name"`
// 	Email    string    `gorm:"column:email; type:varchar(225); not null; unique;"`
// 	IsActive bool      `gorm:"column:is_active; type:boolean; default:false;"`
// }

// func (u *User) TableName() string {
// 	return "go_db_user"
// }
