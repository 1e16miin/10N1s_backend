package room

import (
	"github.com/10n1s-backend/internal/room/repository/database"
)

type Config struct {
	DBConfig database.Config `config:"database"`
}
