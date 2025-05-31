package auth

import (
	"errors"

	"github.com/GuiFernandess7/echoAuthBoilerplate/pkg/logger"
)

type AuthService struct {
	repo AuthRepositoryInterface
}

func NewAuthService(repo AuthRepositoryInterface, logger *logger.Logger) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) Login(dto LoginDTO) (*User, error) {
	user, err := s.repo.FindByEmail(dto.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := CheckPasswordHash(dto.Password, user.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *AuthService) Signup(dto RegisterDTO) (*User, error) {
	_, err := s.repo.FindByEmail(dto.Email)
	if err == nil {
		return nil, errors.New("email already registered")
	}
	if err.Error() != "user not found" {
		return nil, err
	}

	if err := ValidateEmail(dto.Email); err != nil {
		return nil, err
	}
	if err := ValidatePassword(dto.Password); err != nil {
		return nil, err
	}

	passwordHashed, err := HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Email:    dto.Email,
		Password: passwordHashed,
		Name:     dto.Name,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
