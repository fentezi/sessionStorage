package repostirories

import (
	"context"
	"time"

	"github.com/fentezi/session-auth/config"
)

type Redis struct {
}

var ctx = context.Background()

func (r *Redis) CreateSession(uuid string, serializedData uint) error {
	err := config.RDB.Set(ctx, uuid, serializedData, time.Duration(10*time.Minute)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) GetSession(uuid string) (string, error) {
	val, err := config.RDB.Get(ctx, uuid).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *Redis) DeleteSession(uuid string) error {
	err := config.RDB.Del(ctx, uuid).Err()
	if err != nil {
		return err
	}
	return nil
}
