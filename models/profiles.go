package models

import "github.com/jinzhu/gorm"

type Profiles struct {
	gorm.Model
	Name      string `json:"name" gorm:"not null;" sql:"unique"`
	GroupName string `json:"group" gorm:"not null;"`
}
