package services

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"oly-server/internal/models"
)

type AuthService struct {
	db *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{db: db}
}

type SignUpInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	User    *models.User   `json:"user,omitempty"`
	Token   string         `json:"token,omitempty"`
}

func generateID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (s *AuthService) Register(input SignUpInput) (*AuthResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	id, err := generateID()
	if err != nil {
		return nil, err
	}

	var newUser models.User
	err = s.db.QueryRow(
		`INSERT INTO "user" (id, username, email, password) VALUES ($1, $2, $3, $4) RETURNING id, username, email, created_at`,
		id, input.Username, input.Email, string(hashedPassword),
	).Scan(&newUser.ID, &newUser.Username, &newUser.Email, &newUser.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Success: true,
		Message: "Пользователь успешно зарегистрирован",
		User:    &newUser,
	}, nil
}

func (s *AuthService) Login(input SignInInput) (*AuthResponse, error) {
	var user models.User
	err := s.db.QueryRow(
		`SELECT id, username, email, password, created_at FROM "user" WHERE email = $1 LIMIT 1`,
		input.Email,
	).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)

	if err == sql.ErrNoRows {
		return &AuthResponse{
			Success: false,
			Message: "Неверный email или password",
		}, nil
	}
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return &AuthResponse{
			Success: false,
			Message: "Неверный email или password",
		}, nil
	}

	return &AuthResponse{
		Success: true,
		Message: "Вы успешно вошли в систему",
		User: &models.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

var ErrInvalidCredentials = errors.New("invalid credentials")
