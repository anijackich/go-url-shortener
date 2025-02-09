package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	ShortCode string `gorm:"uniqueIndex;not null"`
	URL       string `gorm:"uniqueIndex;not null"`
}
