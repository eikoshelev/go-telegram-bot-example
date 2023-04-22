package bot

import (
	commonanswers "github.com/eikoshelev/go-telegram-bot-template/internal/messages/common-answers"
	"github.com/eikoshelev/go-telegram-bot-template/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (b *Bot) UpdateRouter(upd tgbotapi.Update) {
	updLocal := model.DecodeToLocal(upd)
	if msg := upd.Message; msg != nil {
		if msg.IsCommand() {
			b.SendMessage(b.CommandsHandler(upd.Message.Command(), updLocal))
		} else {
			b.SendMessage(b.MessageHandler(upd))
		}
	}
	if cq := upd.CallbackQuery; cq != nil {
		b.SendMessage(b.CallbacksHandler(updLocal))
	}
}

func (b *Bot) CommandsHandler(command string, updLocal *model.UpdateLocal) tgbotapi.Chattable {
	return commonanswers.UnknownCommand().BuildBotMessage(int64(updLocal.TelegramChatID))
}

func (b *Bot) MessageHandler(upd tgbotapi.Update) tgbotapi.Chattable {
	updLocal := model.DecodeToLocal(upd)
	/*
		your message processing logic should be here
		return <message>
	*/
	return commonanswers.UnknownMessage().BuildBotMessage(int64(updLocal.TelegramChatID))
}

func (b *Bot) CallbacksHandler(updLocal *model.UpdateLocal) tgbotapi.Chattable {
	cData := updLocal.CallbackData
	replyMessage, err := b.Flow.Handle(&cData, updLocal)
	if err != nil {
		b.Logger.Error("error", zap.String("reason", err.Error()))
		return commonanswers.UnknownError().BuildBotMessage(int64(updLocal.TelegramChatID))
	}
	return replyMessage
}
