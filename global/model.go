package global

import (
	"time"

	"gorm.io/gorm"
)

type MAY_MODEL struct {
	ID        uint      `gorm:"primarykey"` // 主键 ID
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 更新时间
}
