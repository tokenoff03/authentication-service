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

	user, err := i.authService.Login(ctx, &model.Login{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user (%v) from db: %v", req.Email, err)
	}

	if !utils.VerifyPassword(req.Password, user.Password) {
		return nil, status.Errorf(codes.Canceled, "password is incorrect")
	}

	refreshToken, err := utils.GenerateToken(model.UserInfo{
		Email: user.Email,
		Role:  user.Role,
	},
		[]byte(i.tokenConfig.RefreshSecretKey()),
		i.tokenConfig.RefreshTTL(),
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	return &auth_v1.LoginResponse{RefreshToken: refreshToken}, nil

}
