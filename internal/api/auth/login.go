package auth

import (
	"authentication-service/internal/model"
	"authentication-service/internal/utils"
	"authentication-service/pkg/auth_v1"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *AuthImplementation) Login(ctx context.Context, req *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {
	//TODO:
	//Прикрутить редис или напрямую идти в базу
	//Сверяем хэш(хэшировать пароль)

	refreshToken, err := utils.GenerateToken(model.UserInfo{
		Email: req.GetEmail(),
		Role:  "admin",
	},
		[]byte(i.tokenConfig.RefreshSecretKey()),
		i.tokenConfig.RefreshTTL(),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	return &auth_v1.LoginResponse{RefreshToken: refreshToken}, nil

}
