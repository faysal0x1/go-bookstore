package repository

import (
	"context"

	"github.com/faysal0x1/go-bookstore/pkg/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(ctx context.Context, book *models.Book) error
	GetAll(ctx context.Context) ([]models.Book, error)
	GetByID(ctx context.Context, id int64) (*models.Book, error)
	Update(ctx context.Context, book *models.Book) error
	Delete(ctx context.Context, id int64) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) Create(ctx context.Context, book *models.Book) error {
	return r.db.WithContext(ctx).Create(book).Error
}

func (r *bookRepository) GetAll(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	err := r.db.WithContext(ctx).Find(&books).Error
	return books, err
}

func (r *bookRepository) GetByID(ctx context.Context, id int64) (*models.Book, error) {
	var book models.Book
	err := r.db.WithContext(ctx).First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) Update(ctx context.Context, book *models.Book) error {
	return r.db.WithContext(ctx).Save(book).Error
}

func (r *bookRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&models.Book{}, id).Error
}
