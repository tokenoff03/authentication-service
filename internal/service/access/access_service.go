package access

import (
	"github.com/tokenoff03/authentication-service/internal/repository"
	"github.com/tokenoff03/authentication-service/internal/service"
	"context"
)

type accessService struct {
	accessRepository repository.AccessRepository
}

func NewAccessService(accessRepository repository.AccessRepository) service.AccessService {
	return &accessService{
		accessRepository: accessRepository,
	}
}

func (s *accessService) Check(ctx context.Context, endpointAdress string) error {
	return nil
}
