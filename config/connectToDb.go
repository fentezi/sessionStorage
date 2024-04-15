package config

import (
	"context"
	"os"

	"github.com/fentezi/session-auth/internal/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client

func MustConnectToSQLDb() {
	var err error
	dsn := os.Getenv("DB_SQL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err = migrateDb(); err != nil {
		panic(err)
	}
}

func migrateDb() error {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	return nil
}

func MustConnectToRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	if err := RDB.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
