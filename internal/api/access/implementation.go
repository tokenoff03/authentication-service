package access

import (
	"authentication-service/internal/config"
	"authentication-service/internal/service"
	"authentication-service/pkg/access_v1"
)

type AccessImplementation struct {
	access_v1.UnimplementedAccessV1Server
	accessService service.AccessService
	tokenConfig   config.TokenConfig
}

func NewAccessImplementation(accessService service.AccessService, tokenConfig config.TokenConfig) *AccessImplementation {
	return &AccessImplementation{accessService: accessService, tokenConfig: tokenConfig}
}
