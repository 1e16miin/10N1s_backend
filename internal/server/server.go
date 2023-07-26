package server

import (
	"fmt"

	"github.com/10n1s-backend/internal/config"
	"github.com/10n1s-backend/pkg/cache"
	"github.com/10n1s-backend/pkg/database"
	"github.com/10n1s-backend/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Server struct {
	echo   *echo.Echo
	cfg    *config.Config
	db     *gorm.DB
	cache  *redis.Client
	logger logger.Logger
}

func NewServer(cfg *config.Config, logger logger.Logger) (*Server, error) {
	echo := echo.New()

	logger.Info("start init db")
	db, err := database.NewDatabase(cfg.DatabaseConfig)
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}
	logger.Info("start init cache")
	cache, err := cache.NewCache(cfg.CacheConfig)
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	return &Server{echo: echo, cfg: cfg, db: db, cache: cache, logger: logger}, nil
}

func (s *Server) Run() error {
	if err := s.MapHandlers(); err != nil {
		s.logger.Fatal(err)
		return err
	}

	return s.echo.Start(fmt.Sprintf(":%s", s.cfg.ServerConfig.Port))
}
