package main

import (
	"context"
	"flag"

	"github.com/10n1s-backend/cmd/auth"
	"github.com/10n1s-backend/cmd/config"
	"github.com/10n1s-backend/cmd/controller"
	"github.com/10n1s-backend/cmd/game"
	"github.com/10n1s-backend/cmd/group"
	"github.com/10n1s-backend/cmd/rank"
	"github.com/10n1s-backend/cmd/repository"
	"github.com/10n1s-backend/cmd/route"
	"github.com/10n1s-backend/cmd/user"
	"github.com/10n1s-backend/pkg/logger"
)

func main() {
	ctx := context.Background()
	configFilePath := flag.String("config", "/etc/10n1s/", "config for 10n1s default path is '/etc/10n1s'")
	flag.Parse()
	cfg := config.Get(*configFilePath)

	logger := logger.GetLogger()

	db, err := repository.NewGormHelper(ctx, cfg.RepositoryConfig)
	if err != nil {
		logger.Fatal("repository init error : " + err.Error())
	}

	authHandler, err := auth.NewAuthHandler(ctx, db)
	if err != nil {
		logger.Fatal("authHandler init error : " + err.Error())
	}

	gameHandler, err := game.NewGameHandler(ctx)
	if err != nil {
		logger.Fatal("gameHandler init error : " + err.Error())
	}

	groupHandler, err := group.NewGroupHandler(ctx)
	if err != nil {
		logger.Fatal("groupHandler init error : " + err.Error())
	}

	rankHandler, err := rank.NewRankHandler(ctx)
	if err != nil {
		logger.Fatal("rankHandler init error : " + err.Error())
	}

	userHandler, err := user.NewUserHandler(ctx)
	if err != nil {
		logger.Fatal("userHandler init error : " + err.Error())
	}

	controllerHandler, err := controller.NewControllerHandler(ctx, authHandler, gameHandler, groupHandler, rankHandler, userHandler)
	if err != nil {
		logger.Fatal("controllerHandler init error : " + err.Error())
	}

	router, _ := route.NewRouter(ctx, cfg.RouteConfig, controllerHandler)
	if err := router.Start(); err != nil {
		logger.Fatal("router start error occurred : " + err.Error())
	}
}
