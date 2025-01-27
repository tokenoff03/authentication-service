package auth

import (
	"authentication-service/internal/model"
	"authentication-service/internal/repository"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/tokenoff03/lib_ad1lek/pkg/db"
)

var (
	psq = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.AuthRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Login(ctx context.Context, login *model.Login) (string, error) {

	return "", nil
}

func (r *repo) GetRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	return "", nil
}

func (r *repo) GetAccessToken(ctx context.Context, accessToken string) (string, error) {
	return "", nil
}
