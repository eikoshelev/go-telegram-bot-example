package model

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// a local object with the necessary data to work on updates coming from Telegram
type UpdateLocal struct {
	TelegramUserID TelegramUserID
	TelegramChatID TelegramChatID
	CallbackData   CallbackData
}

// function to decode the incoming update into the local model
func DecodeToLocal(upd tgbotapi.Update) *UpdateLocal {
	tgUser := upd.SentFrom()
	tgChat := upd.FromChat()
	var cData CallbackData
	if query := upd.CallbackQuery; query != nil {
		cDataBot := CallbackDataBot(upd.CallbackData())
		cData = *cDataBot.Decode()
	}
	return &UpdateLocal{
		TelegramUserID: TelegramUserID(tgUser.ID),
		TelegramChatID: TelegramChatID(tgChat.ID),
		CallbackData:   cData,
	}
}
