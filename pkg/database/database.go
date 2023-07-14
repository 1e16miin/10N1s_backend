package database

import (
	"gorm.io/gorm"
)

type Config struct {
	Engine string      `config:"engine"`
	MySQL  MysqlConfig `config:"mysql"`
}

func NewDatabase(cfg Config) (*gorm.DB, error) {
	switch cfg.Engine {
	case "mysql":
		return initMysqlDB(cfg.MySQL)
	default:
		return initMysqlDB(cfg.MySQL)
	}
}
