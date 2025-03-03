package access

import (
	"context"
	"errors"
	"strings"

	"github.com/tokenoff03/authentication-service/internal/utils"
	"github.com/tokenoff03/authentication-service/pkg/access_v1"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	authPrefix = "Bearer "
)

var accessibleRoles map[string]string

func (i *AccessImplementation) Check(ctx context.Context, req *access_v1.CheckRequest) (*emptypb.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return nil, errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	claims, err := utils.VerifyToken(accessToken, []byte(i.tokenConfig.AccessSecretKey()))
	if err != nil {
		return nil, errors.New("access token is invalid")
	}

	accessibleMap, err := i.accessibleRoles(ctx)
	if err != nil {
		return nil, errors.New("failed to get accessible roles")
	}

	role, ok := accessibleMap[req.GetEndpointAddress()]
	if !ok {
		return &emptypb.Empty{}, nil
	}

	if role == claims.Role {
		return &emptypb.Empty{}, nil
	}

	return nil, errors.New("access denied")
}

func (i *AccessImplementation) accessibleRoles(ctx context.Context) (map[string]string, error) {
	if accessibleRoles == nil {
		accessibleRoles = make(map[string]string)

		// Лезем в базу за данными о доступных ролях для каждого эндпоинта
		// Можно кэшировать данные, чтобы не лезть в базу каждый раз

		// Например, для эндпоинта /note_v1.NoteV1/Get доступна только роль admin
		accessibleRoles["/note_v1.NoteV1/Get"] = "admin"
	}

	return accessibleRoles, nil
}
