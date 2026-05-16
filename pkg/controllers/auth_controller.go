package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/faysal0x1/go-bookstore/pkg/models"
	"github.com/faysal0x1/go-bookstore/pkg/responses"
	"github.com/faysal0x1/go-bookstore/pkg/services"
	"github.com/faysal0x1/go-bookstore/pkg/validators"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := validators.ValidateStruct(&user); err != nil {
		responses.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.Register(r.Context(), &user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses.JSON(w, http.StatusCreated, "User registered successfully", nil)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	token, err := c.service.Login(r.Context(), credentials.Email, credentials.Password)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	responses.JSON(w, http.StatusOK, "Login successful", map[string]string{"token": token})
}
