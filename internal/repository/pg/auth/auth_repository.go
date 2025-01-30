package auth

import (
	"authentication-service/internal/model"
	"authentication-service/internal/repository"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/tokenoff03/lib_ad1lek/pkg/db"
)

const (
	tableName      = `"user"`
	passwordColumn = "password"
	emailColumn    = "email"
	roleColumn     = "role"
)

var (
	psq = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

type repo struct {
	db db.Client
}

func NewAuthRepository(db db.Client) repository.AuthRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Login(ctx context.Context, login *model.Login) (*model.UserInfo, error) {
	builder := psq.Select(passwordColumn, emailColumn, roleColumn).
		From(tableName).
		Where(sq.Eq{emailColumn: login.Email}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "auth_repository.Login",
		QueryRow: query,
	}
	var user model.UserInfo
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
