package services

import (
	"context"
	"errors"
	"time"

	"github.com/faysal0x1/go-bookstore/pkg/config"
	"github.com/faysal0x1/go-bookstore/pkg/models"
	"github.com/faysal0x1/go-bookstore/pkg/repository"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Register(ctx context.Context, user *models.User) error
	Login(ctx context.Context, email, password string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type authService struct {
	repo         repository.AuthRepository
	config       *config.Config
	emailService EmailService
}

func NewAuthService(repo repository.AuthRepository, config *config.Config, emailService EmailService) AuthService {
	return &authService{repo: repo, config: config, emailService: emailService}
}

func (s *authService) Register(ctx context.Context, user *models.User) error {
	// Check if user already exists
	existingUser, _ := s.repo.FindByEmail(ctx, user.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	// Assign default "User" role if not specified
	if user.RoleID == 0 {
		role, err := s.repo.FindRoleByName(ctx, "User")
		if err == nil {
			user.RoleID = role.ID
		}
	}

	err := s.repo.CreateUser(ctx, user)
	if err == nil {
		s.emailService.SendWelcomeEmail(user.Email, user.Name)
	}
	return err
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !user.CheckPassword(password) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role.Name,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(s.config.JWTSecret))
}

func (s *authService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.config.JWTSecret), nil
	})
}
