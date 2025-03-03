package access

import (
	"github.com/tokenoff03/authentication-service/pkg/access_v1"

	"github.com/tokenoff03/authentication-service/internal/config"
	"github.com/tokenoff03/authentication-service/internal/service"
)

type AccessImplementation struct {
	access_v1.UnimplementedAccessV1Server
	accessService service.AccessService
	tokenConfig   config.TokenConfig
}

func NewAccessImplementation(accessService service.AccessService, tokenConfig config.TokenConfig) *AccessImplementation {
	return &AccessImplementation{accessService: accessService, tokenConfig: tokenConfig}
}
