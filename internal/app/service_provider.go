package app

import (
	accessImpl "authentication-service/internal/api/access"
	"authentication-service/internal/api/auth"
	"authentication-service/internal/config"
	"authentication-service/internal/config/env"
	"authentication-service/internal/repository"
	"authentication-service/internal/repository/access"
	authRepo "authentication-service/internal/repository/auth"
	"authentication-service/internal/service"
	accessService "authentication-service/internal/service/access"
	authService "authentication-service/internal/service/auth"
	"context"
	"log"

	"github.com/tokenoff03/lib_ad1lek/pkg/closer"
	"github.com/tokenoff03/lib_ad1lek/pkg/db"
	"github.com/tokenoff03/lib_ad1lek/pkg/db/pg"
)

type serviceProvider struct {
	pgConfig    config.PgConfig
	grpcConfig  config.GRPCConfig
	tokenConfig config.TokenConfig
	dbClient    db.Client

	authRepository   repository.AuthRepository
	accessRepository repository.AccessRepository
	authService      service.AuthService
	accessService    service.AccessService
	authImpl         *auth.AuthImplementation
	accessImpl       *accessImpl.AccessImplementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PgConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %v", err)
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) GRCPConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err)
		}
		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) TokenConfig() config.TokenConfig {
	if s.tokenConfig == nil {
		cfg, err := env.NewTokenConfig()
		if err != nil {
			log.Fatalf("failed to get token config: %v", err)
		}
		s.tokenConfig = cfg
	}

	return s.tokenConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}
		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %v", err)
		}

		closer.Add(cl.Close)
		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = authRepo.NewAuthRepository(s.DBClient(ctx))
	}

	return s.authRepository
}

func (s *serviceProvider) AccessRepository(ctx context.Context) repository.AccessRepository {
	if s.accessRepository == nil {
		s.accessRepository = access.NewAccessRepository(s.DBClient(ctx))

	}

	return s.accessRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewAuthService(s.AuthRepository(ctx))
	}

	return s.authService
}

func (s *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewAccessService(s.AccessRepository(ctx))
	}

	return s.accessService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.AuthImplementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewAuthImplementation(s.AuthService(ctx), s.TokenConfig())
	}

	return s.authImpl
}

func (s *serviceProvider) AccessImpl(ctx context.Context) *accessImpl.AccessImplementation {
	if s.accessImpl == nil {
		s.accessImpl = accessImpl.NewAccessImplementation(s.AccessService(ctx), s.TokenConfig())
	}

	return s.accessImpl
}
