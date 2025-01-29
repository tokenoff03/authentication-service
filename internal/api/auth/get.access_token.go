package auth

import (
	"authentication-service/internal/model"
	"authentication-service/internal/utils"
	"authentication-service/pkg/auth_v1"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *AuthImplementation) GetAccessToken(ctx context.Context, req *auth_v1.GetAccessTokenRequest) (*auth_v1.GetAccessTokenResponse, error) {
	claims, err := utils.VerifyToken(req.GetRefreshToken(), []byte(i.tokenConfig.RefreshSecretKey()))
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "invalid refresh token")
	}

	//Пойди в базу или в кеш за доп данными пользователя

	accessToken, err := utils.GenerateToken(model.UserInfo{
		Email: claims.Email,
		Role:  claims.Role,
	},
		[]byte(i.tokenConfig.RefreshSecretKey()),
		i.tokenConfig.RefreshTTL(),
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	return &auth_v1.GetAccessTokenResponse{AccessToken: accessToken}, nil

}
