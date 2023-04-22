package model

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type (
	CommandKey    string
	CommandEntity struct {
		Key  CommandKey
		Name string
	}
	BotCommands struct {
		Config   tgbotapi.SetMyCommandsConfig
		Commands map[CommandKey]CommandEntity
	}
)

const (
	// consider "start" command only as a command to launch the bot, we do not use it when configuring the menu
	StartCommand   = CommandKey("start")
	ProfileCommand = CommandKey("profile")
	/*
		add your commands here
	*/
)
