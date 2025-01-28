package access

import (
	"authentication-service/internal/service"
	"authentication-service/pkg/access_v1"
)

type AccessImplementation struct {
	access_v1.UnimplementedAccessV1Server
	accessService service.AccessService
}

func NewAccessImplementation(accessService service.AccessService) *AccessImplementation {
	return &AccessImplementation{accessService: accessService}
}
