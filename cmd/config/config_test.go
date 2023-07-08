package config

import (
	"testing"

	"github.com/10n1s-backend/cmd/auth"
	"github.com/10n1s-backend/cmd/game"
	"github.com/10n1s-backend/cmd/group"
	"github.com/10n1s-backend/cmd/rank"
	"github.com/10n1s-backend/cmd/route"
	"github.com/10n1s-backend/cmd/user"
	"github.com/10n1s-backend/pkg/database"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	tennisConfig := &TennisConfig{
		RouteConfig: route.Config{
			Type: "echo",
			Echo: route.EchoConfig{
				Port: "1234",
			},
		},
		DatabaseConfig: database.Config{
			Engine: "mysql",
			MySQL: database.MysqlConfig{
				Port:                "3306",
				User:                "root",
				PassWd:              "asdf",
				EndPoint:            "127.0.0.1",
				Database:            "tennis",
				MaxIdleConnections:  100,
				MaxOpenConnections:  100,
				ConnMaxIdleTime:     5,
				ConnMaxLifetime:     5,
				QueryLogModeEnabled: true,
			},
		},
		AuthConfig:  auth.Config{},
		GameConfig:  game.Config{},
		GroupConfig: group.Config{},
		RankConfig:  rank.Config{},
		UserConfig:  user.Config{},
	}

	config := Get("")
	assert.Equal(t, tennisConfig, config)
}
