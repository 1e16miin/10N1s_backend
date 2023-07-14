package main

import (
	"flag"

	"github.com/10n1s-backend/internal/config"
	"github.com/10n1s-backend/internal/server"
	"github.com/10n1s-backend/pkg/logger"
)

func main() {
	configFilePath := flag.String("config", "/etc/10n1s/", "config for 10n1s default path is '/etc/10n1s'")
	flag.Parse()
	cfg := config.Get(*configFilePath)

	logger := logger.NewZapLogger()
	logger.InitLogger()

	server, err := server.NewServer(cfg, logger)
	if err != nil {
		logger.Fatal(err)
		return
	}

	if err := server.Run(); err != nil {
		logger.Fatal("router start error occurred : " + err.Error())
	}
}
