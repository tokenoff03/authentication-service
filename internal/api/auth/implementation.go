package auth

import (
	"authentication-service/internal/config"
	"authentication-service/internal/service"
	"authentication-service/pkg/auth_v1"
)

type AuthImplementation struct {
	auth_v1.UnimplementedAuthV1Server
	authService service.AuthService
	tokenConfig config.TokenConfig
}

func NewAuthImplementation(authService service.AuthService, tokenConfig config.TokenConfig) *AuthImplementation {
	return &AuthImplementation{authService: authService, tokenConfig: tokenConfig}
}
