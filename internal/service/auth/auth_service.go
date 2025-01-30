package service

import (
	"authentication-service/internal/model"
	"authentication-service/internal/repository"
	"authentication-service/internal/service"
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
