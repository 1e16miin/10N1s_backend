package repository

import (
	"context"

	"gorm.io/gorm"
)

var gormDB *gorm.DB

type Config struct {
	Engine string      `config:"engine"`
	MySQL  MysqlConfig `config:"mysql"`
}

func NewGormHelper(ctx context.Context, config Config) (*gorm.DB, error) {
	var err error

	if gormDB != nil {
		switch config.Engine {
		case "mysql":
			gormDB, err = initMysqlDB(config.MySQL)
		default:
			gormDB, err = initMysqlDB(config.MySQL)
		}
		if err != nil {
			return nil, err
		}
		return gormDB, nil
	}
	return nil, nil
}
