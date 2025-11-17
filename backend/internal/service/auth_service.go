package service

import (
	"errors"

	"github.com/Shihasz/gophiway/internal/config"
	"github.com/Shihasz/gophiway/internal/models"
	"github.com/Shihasz/gophiway/internal/repository"
	"github.com/Shihasz/gophiway/pkg/crypto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserNotFound       = errors.New("user not found")
)

type AuthService struct {
	userRepo *repository.UserRepository
	cfg      *config.Config
}

func NewAuthService(userRepo *repository.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

// RegisterRequest represents a user registration request
type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

// LoginRequest represents a user login request
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// AuthResponse represents an authentication response
type AuthResponse struct {
	User         *UserResponse `json:"user"`
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
}

// UserResponse represents a user response
type UserResponse struct {
	ID            uuid.UUID `json:"id"`
	Email         string    `json:"email"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Role          string    `json:"role"`
	EmailVerified bool      `json:"email_verified"`
}

// Register registers a new user
func (s *AuthService) Register(req *RegisterRequest) (*AuthResponse, error) {
	// Check if email already exists
	exists, err := s.userRepo.EmailExists(req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrEmailAlreadyExists
	}

	// Hash password
	hashedPassword, err := crypto.HashPassword(req.Password, s.cfg.BcryptCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		Email:        req.Email,
		PasswordHash: hashedPassword,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Role:         "customer",
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Generate tokens
	accessToken, err := crypto.GenerateToken(user.ID, user.Email, user.Role, s.cfg.JWTSecret, s.cfg.JWTExpiration)
	if err != nil {
		return nil, err
	}

	refreshToken, err := crypto.GenerateToken(user.ID, user.Email, user.Role, s.cfg.JWTRefreshSecret, s.cfg.JWTRefreshExpiration)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		User: &UserResponse{
			ID:            user.ID,
			Email:         user.Email,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			Role:          user.Role,
			EmailVerified: user.EmailVerified,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Login authenticates a user
func (s *AuthService) Login(req *LoginRequest) (*AuthResponse, error) {
	// Get user by email
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	// Check password
	if !crypto.CheckPassword(req.Password, user.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	// Generate tokens
	accessToken, err := crypto.GenerateToken(user.ID, user.Email, user.Role, s.cfg.JWTSecret, s.cfg.JWTExpiration)
	if err != nil {
		return nil, err
	}

	refreshToken, err := crypto.GenerateToken(user.ID, user.Email, user.Role, s.cfg.JWTRefreshSecret, s.cfg.JWTRefreshExpiration)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		User: &UserResponse{
			ID:            user.ID,
			Email:         user.Email,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			Role:          user.Role,
			EmailVerified: user.EmailVerified,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// RefreshToken refreshes an access token
func (s *AuthService) RefreshToken(refreshToken string) (*AuthResponse, error) {
	// Validate refresh token
	claims, err := crypto.ValidateToken(refreshToken, s.cfg.JWTRefreshSecret)
	if err != nil {
		return nil, err
	}

	// Get user
	user, err := s.userRepo.GetByID(claims.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// Generate new tokens
	accessToken, err := crypto.GenerateToken(user.ID, user.Email, user.Role, s.cfg.JWTSecret, s.cfg.JWTExpiration)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := crypto.GenerateToken(user.ID, user.Email, user.Role, s.cfg.JWTRefreshSecret, s.cfg.JWTRefreshExpiration)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		User: &UserResponse{
			ID:            user.ID,
			Email:         user.Email,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			Role:          user.Role,
			EmailVerified: user.EmailVerified,
		},
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

// GetCurrentUser gets the current authenticated user
func (s *AuthService) GetCurrentUser(userID uuid.UUID) (*UserResponse, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &UserResponse{
		ID:            user.ID,
		Email:         user.Email,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Role:          user.Role,
		EmailVerified: user.EmailVerified,
	}, nil
}
