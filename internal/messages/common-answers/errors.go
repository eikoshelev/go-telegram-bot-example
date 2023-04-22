package commonanswers

import "github.com/eikoshelev/go-telegram-bot-template/internal/model"

func UnknownError() model.Message {
	return model.Message{
		Text: "Sorry something went wrong :(\nPlease try again later.",
	}
}
