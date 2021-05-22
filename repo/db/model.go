package db

import "gorm.io/gorm"

type IPAddress struct {
	gorm.Model
	UUID         string `gorm:"not null"`
	ResponseCode string `gorm:"not null"`
	IPAddress    string `gorm:"index;not null"`
}
