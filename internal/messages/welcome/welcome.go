package welcome

import (
	"github.com/eikoshelev/go-telegram-bot-template/internal/model"
)

func Welcome() model.Message {
	return model.Message{
		Text: "Hi! This is a test hello message :)\nYou can create a profile using the button below.",
		Buttons: []model.Button{
			{
				Name: "Create profile",
				CallbackData: model.CallbackData{
					CommandKey: model.StartCommand,
					Case:       "welcome",
					Step:       0,
					Payload:    "start",
				},
			},
		},
	}
}

func CreateProfile() model.Message {
	return model.Message{
		Text: "Profile created successfully!",
	}
}

func ProfileAlreadyExists() model.Message {
	return model.Message{
		Text: "Profile already exists!",
	}
}
