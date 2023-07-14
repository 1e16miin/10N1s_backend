package database

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConfig struct {
	Port                string        `config:"port"`
	User                string        `config:"user"`
	PassWd              string        `config:"passwd"`
	EndPoint            string        `config:"endpoint"`
	Database            string        `config:"database"`
	MaxIdleConnections  int           `config:"maxIdleConnections"`
	MaxOpenConnections  int           `config:"maxOpenConnections"`
	ConnMaxIdleTime     time.Duration `config:"connMaxIdleTime"`
	ConnMaxLifetime     time.Duration `config:"connMaxLifeTime"`
	QueryLogModeEnabled bool          `config:"queryLogMode"`
}

func initMysqlDB(config MysqlConfig) (*gorm.DB, error) {
	db, err := mysqlCreateDBConnectionPool(config)
	if err != nil {
		return nil, err
	}
	return mysqlCreateGormDB(config, db)
}

func mysqlCreateDBConnectionPool(config MysqlConfig) (*sql.DB, error) {
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC&autocommit=false&tls=preferred",
		config.User, config.PassWd, config.EndPoint, config.Port, config.Database)
	sqlDB, err := sql.Open("mysql", connectString)
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(config.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(config.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(config.ConnMaxIdleTime)

	return sqlDB, nil
}

func mysqlCreateGormDB(config MysqlConfig, db *sql.DB) (*gorm.DB, error) {
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}))
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}
