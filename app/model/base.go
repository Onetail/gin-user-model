package model

import (
	"time"
)

type BaseModel struct {
	ID        string    `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time `gorm:"index:idx_time"`
	UpdatedAt time.Time `gorm:"index:idx_time"`
}