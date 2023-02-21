package global

import (
	"gorm.io/gorm"
	"time"
)

type GVA_MODEL struct {
	// 主键 ID
	ID uint `gorm:"primarykey"`
	// 创建时间
	CreatedAt time.Time
	// 更新时间
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
