package seeders

import (
	"log"

	"github.com/faysal0x1/go-bookstore/pkg/models"
	"gorm.io/gorm"
)

type BookSeeder struct{}

func (s *BookSeeder) Run(db *gorm.DB) error {
	var count int64
	db.Model(&models.Book{}).Count(&count)
	if count > 0 {
		return nil
	}

	books := []models.Book{
		{Name: "The Go Programming Language", Author: "Alan A. A. Donovan", Publication: "Addison-Wesley"},
		{Name: "Clean Code", Author: "Robert C. Martin", Publication: "Prentice Hall"},
		{Name: "Refactoring", Author: "Martin Fowler", Publication: "Addison-Wesley"},
	}

	for _, book := range books {
		if err := db.Create(&book).Error; err != nil {
			return err
		}
	}

	log.Println("BookSeeder: Database seeded successfully")
	return nil
}
