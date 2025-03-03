package redis_repository

import (
	"github.com/tokenoff03/authentication-service/internal/model"
	"github.com/tokenoff03/authentication-service/internal/repository"
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/tokenoff03/lib_ad1lek/pkg/cache"
)

const tableKey = "users"

type userRepository struct {
	redisClient cache.RedisClient
}

func NewUserRepository(redisClient cache.RedisClient) repository.UserCacheRepository {
	return &userRepository{
		redisClient: redisClient,
	}
}

func (r *userRepository) GetUser(ctx context.Context, email string) (*model.UserInfo, error) {
	result, err := r.redisClient.HGet(ctx, tableKey, email)
	if err != nil {
		return nil, err
	}

	var retrievedUser model.UserInfo
	err = json.Unmarshal([]byte(result), &retrievedUser)
	if err != nil {
		return nil, errors.Wrap(err, "JSON Unmarshal error")
	}

	return &retrievedUser, nil
}

func (r *userRepository) SetUser(ctx context.Context, user *model.UserInfo) error {
	userJson, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "JSON Marshal error")
	}

	return r.redisClient.HashSet(ctx, tableKey, map[string]interface{}{user.Email: userJson})
}
