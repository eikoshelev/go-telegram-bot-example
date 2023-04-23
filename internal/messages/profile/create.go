package profile

import "github.com/eikoshelev/go-telegram-bot-example/internal/model"

func Create() model.Message {
	return model.Message{
		Text: "This is a demo message from a bot built using the example taken from this repository - https://github.com/eikoshelev/go-telegram-bot-example.\n\nAs a demonstration, by clicking the button below, you can create a minimal user profile with real data from Telegram.",
		Buttons: []model.Button{
			{
				Name: "Create profile",
				CallbackData: model.CallbackData{
					CommandKey: model.ProfileCommand,
					Case:       "create",
					Step:       0,
					Payload:    "",
				},
			},
		},
	}
}
