package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/faysal0x1/go-bookstore/pkg/models"
	"github.com/faysal0x1/go-bookstore/pkg/responses"
	"github.com/faysal0x1/go-bookstore/pkg/services"
	"github.com/faysal0x1/go-bookstore/pkg/utils"
	"github.com/faysal0x1/go-bookstore/pkg/validators"
	"github.com/gorilla/mux"
)

type BookController struct {
	service services.BookService
}

func NewBookController(service services.BookService) *BookController {
	return &BookController{service: service}
}

func (c *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.service.GetAllBooks(r.Context())
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Could not fetch books")
		return
	}
	responses.JSON(w, http.StatusOK, "Books fetched successfully", books)
}

func (c *BookController) GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["bookId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book, err := c.service.GetBookByID(r.Context(), id)
	if err != nil {
		responses.Error(w, http.StatusNotFound, "Book not found")
		return
	}
	responses.JSON(w, http.StatusOK, "Book fetched successfully", book)
}

func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	// Parse Multipart Form
	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10MB max
		responses.Error(w, http.StatusBadRequest, "Failed to parse multipart form")
		return
	}

	var book models.Book
	book.Name = r.FormValue("name")
	book.Author = r.FormValue("author")
	book.Publication = r.FormValue("publication")

	// Validate basic fields
	if err := validators.ValidateStruct(&book); err != nil {
		responses.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Handle File Upload
	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		path, err := utils.UploadFile(file, header)
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, "Failed to upload image")
			return
		}
		book.Image = path
	}

	if err := c.service.CreateBook(r.Context(), &book); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Could not create book")
		return
	}
	responses.JSON(w, http.StatusCreated, "Book created successfully", book)
}

func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["bookId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	var updateData models.Book
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	book, err := c.service.UpdateBook(r.Context(), id, &updateData)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Could not update book")
		return
	}
	responses.JSON(w, http.StatusOK, "Book updated successfully", book)
}

func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["bookId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	if err := c.service.DeleteBook(r.Context(), id); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Could not delete book")
		return
	}
	responses.JSON(w, http.StatusOK, "Book deleted successfully", nil)
}
