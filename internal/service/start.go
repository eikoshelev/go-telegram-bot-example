package service

import (
	"github.com/eikoshelev/go-telegram-bot-template/internal/model"
	"github.com/eikoshelev/go-telegram-bot-template/internal/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Start struct {
	Start *repository.Start
}

func (s *Start) StartHandler(updLocal *model.UpdateLocal) (tgbotapi.Chattable, error) {
	/*
		there can be any business logic necessary for the further functioning of the application,
		for example, various checks and data preparation
	*/
	return nil, nil
}
