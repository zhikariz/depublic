package database

import (
	"fmt"

	"github.com/zhikariz/depublic/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToPostgres(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.Postgres.Host,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
		cfg.Postgres.Port)

	var log logger.Interface

	if cfg.Env == "dev" {
		log = logger.Default.LogMode(logger.Info)
	} else {
		log = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
