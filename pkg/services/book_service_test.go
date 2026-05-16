package services

import (
	"context"
	"testing"

	"github.com/faysal0x1/go-bookstore/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository is a mock of BookRepository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(ctx context.Context, book *models.Book) error {
	args := m.Called(ctx, book)
	return args.Error(0)
}

func (m *MockRepository) GetAll(ctx context.Context) ([]models.Book, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Book), args.Error(1)
}

func (m *MockRepository) GetByID(ctx context.Context, id int64) (*models.Book, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Book), args.Error(1)
}

func (m *MockRepository) Update(ctx context.Context, book *models.Book) error {
	args := m.Called(ctx, book)
	return args.Error(0)
}

func (m *MockRepository) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestGetAllBooks(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewBookService(mockRepo)
	ctx := context.Background()

	expectedBooks := []models.Book{
		{Name: "Book 1", Author: "Author 1"},
		{Name: "Book 2", Author: "Author 2"},
	}

	mockRepo.On("GetAll", ctx).Return(expectedBooks, nil)

	books, err := service.GetAllBooks(ctx)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(books))
	assert.Equal(t, "Book 1", books[0].Name)
	mockRepo.AssertExpectations(t)
}
