package service

import (
	"github.com/tokenoff03/authentication-service/internal/model"

	"context"
)

type AuthService interface {
	Login(ctx context.Context, login *model.Login) (*model.UserInfo, error)
}

type AccessService interface {
	Check(ctx context.Context, endpointAdress string) error
}

type UserCacheService interface {
	GetUser(ctx context.Context, email string) (*model.UserInfo, error)
	SetUser(ctx context.Context, user *model.UserInfo) error
}
