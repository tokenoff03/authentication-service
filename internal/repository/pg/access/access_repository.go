package access

import (
	"github.com/tokenoff03/authentication-service/internal/repository"
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

func NewAccessRepository(db db.Client) repository.AccessRepository {
	return &repo{
		db: db,
	}
}

func (s *repo) Check(ctx context.Context, endpointAdress string) error {
	return nil
}
