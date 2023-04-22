package main

import (
	"flag"

	"github.com/eikoshelev/go-telegram-bot-template/internal/bot"
	"github.com/eikoshelev/go-telegram-bot-template/internal/config"
	"github.com/eikoshelev/go-telegram-bot-template/internal/database"
	"github.com/eikoshelev/go-telegram-bot-template/internal/flow"
	"github.com/eikoshelev/go-telegram-bot-template/internal/logger"
	"github.com/eikoshelev/go-telegram-bot-template/internal/repository"
	"github.com/eikoshelev/go-telegram-bot-template/internal/service"

	"go.uber.org/zap"
)

func main() {
	configPath := flag.String("c", "./cmd/go-telegram-bot-template/config.yaml", "path to go-telegram-bot-template config")
	flag.Parse()

	logger := logger.GetLogger()

	cfg := &config.Config{}

	err := config.GetConfiguration(*configPath, cfg)
	if err != nil {
		logger.Fatal("failed get configuration", zap.String("reason", err.Error()))
	}

	logger.Info("configured", zap.Any("config", cfg))

	db, err := database.NewPostgresDB(cfg.Database)
	if err != nil {
		logger.Fatal("failed connect to DB", zap.String("reason", err.Error()))
	}
	logger.Info("success connected to database")

	err = database.Migrate(&cfg.Migration, logger)
	if err != nil {
		logger.Fatal("can't run db migrations", zap.String("reason", err.Error()))
	}

	repo := repository.Init(db)
	svc := service.Init(repo)
	flows := flow.Init(svc)

	flows.ValidateCallbacksDataSize(&logger)

	bot := bot.Init(cfg, &logger, flows, svc, repo)

	bot.Run()
}
