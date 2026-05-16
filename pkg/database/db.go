package database

import (
	"fmt"
	"log"

	"github.com/faysal0x1/go-bookstore/pkg/config"
	"github.com/faysal0x1/go-bookstore/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
	return db
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Book{}, &models.User{}, &models.Role{}, &models.Permission{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Database migration completed")
}

func DropTables(db *gorm.DB) {
	log.Println("Dropping all tables...")
	err := db.Migrator().DropTable(&models.Book{}, &models.User{}, &models.Role{}, &models.Permission{}, "role_permissions")
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}
	log.Println("Tables dropped successfully")
}

