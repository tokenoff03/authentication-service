package auth

import (
	"authentication-service/internal/service"
	"authentication-service/pkg/access_v1"
	"authentication-service/pkg/auth_v1"
)

type Implementation struct {
	auth_v1.UnimplementedAuthV1Server
	access_v1.UnimplementedAccessV1Server
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{authService: authService}
}
