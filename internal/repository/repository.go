package repository

import (
	"authentication-service/internal/model"
	"context"
)

type AuthRepository interface {
	Login(ctx context.Context, login *model.Login) (*model.UserInfo, error)
}

type AccessRepository interface {
	Check(ctx context.Context, endpointAdress string) error
}

type UserCacheRepository interface {
	GetUser(ctx context.Context, email string) (*model.UserInfo, error)
	SetUser(ctx context.Context, user *model.UserInfo) error
}
