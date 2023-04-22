package service

import (
	"time"

	commonanswers "github.com/eikoshelev/go-telegram-bot-template/internal/messages/common-answers"
	"github.com/eikoshelev/go-telegram-bot-template/internal/model"
	"github.com/eikoshelev/go-telegram-bot-template/internal/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type WelcomeService struct {
	Welcome *repository.WelcomeRepo
}

func (ws *WelcomeService) WelcomeHandler(updLocal *model.UpdateLocal) (tgbotapi.Chattable, error) {
	return nil, nil
}

func (ws *WelcomeService) CreateProfile(updLocal *model.UpdateLocal) (tgbotapi.Chattable, error) {
	/*
		below is an example of how you can interact with the db,
		uncomment the call to the repository method and create a user to write information about the user (don't forget to prepare the db)
	*/

	userProfile := &model.UserProfile{
		TelegramUserID: updLocal.TelegramUserID,
		TelegramChatID: updLocal.TelegramChatID,
		CreatedAt:      time.Now(),
	}
	err := ws.Welcome.CreateProfile(userProfile)
	if err != nil {
		reply := commonanswers.UnknownError().BuildBotMessage(int64(updLocal.TelegramChatID))
		return reply, err
	}
	return nil, nil
}
