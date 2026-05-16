package seeders

import (
	"log"

	"github.com/faysal0x1/go-bookstore/pkg/models"
	"gorm.io/gorm"
)

type UserSeeder struct{}

func (s *UserSeeder) Run(db *gorm.DB) error {
	var adminRole models.Role
	if err := db.Where("name = ?", "Admin").First(&adminRole).Error; err != nil {
		return err
	}

	adminUser := models.User{
		Name:     "Admin User",
		Email:    "admin@example.com",
		Password: "adminpassword", // Will be hashed by BeforeCreate hook
		RoleID:   adminRole.ID,
	}

	var existingUser models.User
	if err := db.Where("email = ?", adminUser.Email).First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&adminUser).Error; err != nil {
				return err
			}
			log.Println("UserSeeder: Admin user created")
		} else {
			return err
		}
	}

	return nil
}
