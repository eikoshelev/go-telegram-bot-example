package commonanswers

import "github.com/eikoshelev/go-telegram-bot-example/internal/model"

func UnknownCommand() model.Message {
	return model.Message{
		Text: "Sorry, I don't know such a command :(",
	}
}

func UnknownMessage() model.Message {
	return model.Message{
		Text: "Sorry, I didn't understand this message :(",
	}
}
