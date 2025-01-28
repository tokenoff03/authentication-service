package auth

import (
	"authentication-service/internal/service"
	"authentication-service/pkg/auth_v1"
)

type AuthImplementation struct {
	auth_v1.UnimplementedAuthV1Server
	authService service.AuthService
}

func NewAuthImplementation(authService service.AuthService) *AuthImplementation {
	return &AuthImplementation{authService: authService}
}
