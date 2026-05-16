package services

import (
	"context"

	"github.com/faysal0x1/go-bookstore/pkg/models"
	"github.com/faysal0x1/go-bookstore/pkg/repository"
)

type BookService interface {
	CreateBook(ctx context.Context, book *models.Book) error
	GetAllBooks(ctx context.Context) ([]models.Book, error)
	GetBookByID(ctx context.Context, id int64) (*models.Book, error)
	UpdateBook(ctx context.Context, id int64, updateData *models.Book) (*models.Book, error)
	DeleteBook(ctx context.Context, id int64) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) CreateBook(ctx context.Context, book *models.Book) error {
	return s.repo.Create(ctx, book)
}

func (s *bookService) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	return s.repo.GetAll(ctx)
}

func (s *bookService) GetBookByID(ctx context.Context, id int64) (*models.Book, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *bookService) UpdateBook(ctx context.Context, id int64, updateData *models.Book) (*models.Book, error) {
	book, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if updateData.Name != "" {
		book.Name = updateData.Name
	}
	if updateData.Author != "" {
		book.Author = updateData.Author
	}
	if updateData.Publication != "" {
		book.Publication = updateData.Publication
	}

	err = s.repo.Update(ctx, book)
	return book, err
}

func (s *bookService) DeleteBook(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
