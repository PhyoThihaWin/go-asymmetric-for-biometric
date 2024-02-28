package utils

import (
	"time"

	"gorm.io/gorm"
)

type CustomModel struct {
	ID        uint           `gorm:"id, primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"created_at" json:"-"`
	UpdatedAt time.Time      `gorm:"updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at, index" json:"-"`
}
