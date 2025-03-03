package auth

import (
	"context"

	"github.com/tokenoff03/authentication-service/internal/model"
	"github.com/tokenoff03/authentication-service/internal/utils"
	"github.com/tokenoff03/authentication-service/pkg/auth_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const traceIDKey = "x-trace-id"

func (i *AuthImplementation) Login(ctx context.Context, req *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {

	user, err := i.authService.Login(ctx, &model.Login{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user (%v) from db: %v", req.Email, err)
	}

	if !utils.VerifyPassword(user.Password, req.Password) {
		return nil, status.Errorf(codes.Aborted, "password is incorrect")
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
