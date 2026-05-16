package seeders

import (
	"log"

	"gorm.io/gorm"
)

func RunDatabaseSeeder(db *gorm.DB) {
	log.Println("Running Database Seeder...")

	seeders := []Seeder{
		&RoleSeeder{},
		&UserSeeder{},
		&BookSeeder{},
	}

	for _, seeder := range seeders {
		if err := seeder.Run(db); err != nil {
			log.Printf("Seeder failed: %v", err)
		}
	}

	log.Println("Database Seeding Completed")
}
