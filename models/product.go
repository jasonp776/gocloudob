package models

import (
	"time"
)

type Product struct {
	//gorm.Model
	Id           int
	Name         string
	Category     string
	Price        float32
	Photo        string
	Descriptions string
	CreateTime   time.Time
}
