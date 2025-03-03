package redis_service

import (
	"context"

	"github.com/tokenoff03/authentication-service/internal/model"
	"github.com/tokenoff03/authentication-service/internal/repository"
	"github.com/tokenoff03/authentication-service/internal/service"
)

type userService struct {
	userCacheRepository repository.UserCacheRepository
}

func NewUserCacheService(userCacheRepository repository.UserCacheRepository) service.UserCacheService {
	return &userService{
		userCacheRepository: userCacheRepository,
	}
}

func (s *userService) GetUser(ctx context.Context, email string) (*model.UserInfo, error) {
	return s.userCacheRepository.GetUser(ctx, email)
}

func (s *userService) SetUser(ctx context.Context, user *model.UserInfo) error {
	return s.userCacheRepository.SetUser(ctx, user)
}
