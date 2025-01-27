package auth

import (
	"authentication-service/internal/model"
	"authentication-service/internal/repository"
	"authentication-service/internal/service"
	"context"
)

type serv struct {
	authRepository repository.AuthRepository
}

func NewService(authRepository repository.AuthRepository) service.AuthService {
	return &serv{
		authRepository: authRepository,
	}
}

func (s *serv) Login(ctx context.Context, login *model.Login) (string, error) {
	return "", nil
}

func (s *serv) GetRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	return "", nil
}

func (s *serv) GetAccessToken(ctx context.Context, accessToken string) (string, error) {
	return "", nil
}
