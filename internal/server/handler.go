package server

import (
	"strings"

	"github.com/10n1s-backend/internal/room"
	roomCache "github.com/10n1s-backend/internal/room/repository/cache"
	roomDB "github.com/10n1s-backend/internal/room/repository/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) MapHandlers() error {
	roomRepository, err := roomDB.NewRepository(s.cfg.RoomConfig.DBConfig, s.db)
	if err != nil {
		s.logger.Fatal(err)
	}

	roomRepositoryCache := roomCache.NewRepositoryCache()

	roomSVC := room.NewService(roomRepository, roomRepositoryCache, s.db, s.cache, s.logger)

	room.RegisterHandlers(s.echo, roomSVC, s.logger)

	s.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID},
	}))
	s.echo.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	s.echo.Use(middleware.RequestID())

	s.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))

	s.echo.Use(middleware.Secure())

	return nil
}
