package service

import (
	"github.com/eikoshelev/go-telegram-bot-example/internal/model"
	"github.com/eikoshelev/go-telegram-bot-example/internal/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Profile struct {
	Profile *repository.Profile
}

func (p *Profile) Create(updLocal *model.UpdateLocal) (tgbotapi.Chattable, error) {
	return nil, nil
}
