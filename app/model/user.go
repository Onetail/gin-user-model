package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)


func (model *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4().String())
}

type User struct {
	BaseModel
	Username string `gorm:"type:varchar(128);not null"`
	Password string `gorm:"type:varchar(128);not null"`
	Email    string `gorm:"type:varchar(256)"`
	Level    uint   `gorm:"default:1;not null"`
}
