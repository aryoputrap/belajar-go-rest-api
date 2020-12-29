package service

import (
	"belajar-go-rest-api/config"
	"belajar-go-rest-api/entity"
	"belajar-go-rest-api/repository"
)

// AuthService struct
type AuthService struct {
	userService entity.UserService
	jwtService  entity.JWTService
}

// NewAuthService func
func NewAuthService(repo *repository.Repository) entity.AuthService {
	return &AuthService{
		userService: NewUserService(repo),
		jwtService:  NewJWTService(config.ConfigureJWT()),
	}
}

// Login func
func (a AuthService) Login(userInput *entity.User) (string, error) {
	user, err := a.userService.FindByEmail(userInput.Email)
	if err != nil {
		return "", err
	}

	err = user.IsPasswordValid(userInput.Password)
	if err != nil {
		return "", err
	}

	token, err := a.jwtService.Create(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
