package models

import (
	"github.com/jinzhu/gorm"
)

type Credentials struct {
	gorm.Model
	Password  string `json:"password" gorm:"password"`
	Username  string `gorm:"type:varchar(100);unique_index"`
	Email     string `json:"email" gorm:"email"`
}
