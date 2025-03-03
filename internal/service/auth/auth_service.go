package service

import (
	"github.com/tokenoff03/authentication-service/internal/model"
	"github.com/tokenoff03/authentication-service/internal/repository"
	"github.com/tokenoff03/authentication-service/internal/service"
	"context"
)

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) service.AuthService {
	return &authService{
		authRepository: authRepository,
	}
}

func (s *authService) Login(ctx context.Context, login *model.Login) (*model.UserInfo, error) {
	return s.authRepository.Login(ctx, login)
}
