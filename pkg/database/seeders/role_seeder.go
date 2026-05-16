package seeders

import (
	"log"

	"github.com/faysal0x1/go-bookstore/pkg/models"
	"gorm.io/gorm"
)

type RoleSeeder struct{}

func (s *RoleSeeder) Run(db *gorm.DB) error {
	roles := []models.Role{
		{Name: "Admin"},
		{Name: "User"},
	}

	for _, role := range roles {
		var existingRole models.Role
		if err := db.Where("name = ?", role.Name).First(&existingRole).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&role).Error; err != nil {
					return err
				}
				log.Printf("RoleSeeder: Role %s created", role.Name)
			} else {
				return err
			}
		}
	}
	return nil
}
