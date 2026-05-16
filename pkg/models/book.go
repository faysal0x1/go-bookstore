package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name" validate:"required"`
	Author      string `json:"author" validate:"required"`
	Publication string `json:"publication" validate:"required"`
	Image       string `json:"image"`
}
