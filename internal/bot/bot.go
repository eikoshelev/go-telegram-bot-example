package bot

import (
	"net/http"

	"github.com/eikoshelev/go-telegram-bot-example/internal/config"
	"github.com/eikoshelev/go-telegram-bot-example/internal/model"
	"github.com/eikoshelev/go-telegram-bot-example/internal/repository"
	"github.com/eikoshelev/go-telegram-bot-example/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Bot struct {
	API        *tgbotapi.BotAPI
	Config     *config.Config
	Logger     *zap.Logger
	Flow       model.Flow
	Service    *service.Service
	Repository *repository.Repository
}

func Init(config *config.Config, logger *zap.Logger, flow model.Flow, service *service.Service, repository *repository.Repository) *Bot {
	return &Bot{
		Config:     config,
		Logger:     logger,
		Flow:       flow,
		Service:    service,
		Repository: repository,
	}
}

func (b *Bot) Run() {
	botAPI, err := b.NewBotAPI()
	if err != nil {
		b.Logger.Fatal("failed create new bot api instance", zap.String("error", err.Error()))
	}
	b.API = botAPI

	if err := b.SetWebhook(); err != nil {
		b.Logger.Fatal("failed set webhook", zap.String("error", err.Error()))
	}
	if err := b.SetBotCommands(); err != nil {
		b.Logger.Fatal("failed set bot commands", zap.String("error", err.Error()))
	}

	go b.StartWebhookServer()
	b.Logger.Info("http webhook server started")

	updates := b.API.ListenForWebhook("/" + b.API.Token)
	for update := range updates {
		go b.UpdateRouter(update)
	}
}

func (b *Bot) NewBotAPI() (*tgbotapi.BotAPI, error) {
	botAPI, err := tgbotapi.NewBotAPI(b.Config.Bot.Token)
	if err != nil {
		return nil, err
	}
	b.Logger.Info("authorized success, bot api instance created", zap.String("account", botAPI.Self.UserName))
	return botAPI, nil
}

func (b *Bot) SetWebhook() error {
	webhook, err := tgbotapi.NewWebhook(b.Config.Bot.WebhookLink + b.API.Token)
	if err != nil {
		return err
	}
	_, err = b.API.Request(webhook)
	if err != nil {
		return err
	}
	info, err := b.API.GetWebhookInfo()
	if err != nil {
		return err
	}
	b.Logger.Info("webhook info", zap.Any("webhook", info))
	if info.LastErrorDate != 0 {
		return err
	}
	return nil
}

// configure the bot menu, don't use "start" command, but you can if you want
func (b *Bot) InitBotCommands() tgbotapi.SetMyCommandsConfig {
	commands := []model.CommandEntity{
		{
			Key:  model.ProfileCommand,
			Name: "profile",
		},
		/* implement your commands in the same way
		{
			Key:  model.<...>,
			Name: "...",
		},
		...
		*/
	}
	tgCommands := make([]tgbotapi.BotCommand, 0, len(commands))
	for _, cmd := range commands {
		tgCommands = append(tgCommands, tgbotapi.BotCommand{
			Command:     "/" + string(cmd.Key),
			Description: cmd.Name,
		})
	}
	commandsConfig := tgbotapi.NewSetMyCommands(tgCommands...)
	return commandsConfig
}

func (b *Bot) SetBotCommands() error {
	commandsConfig := b.InitBotCommands()
	_, err := b.API.Request(commandsConfig)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) StartWebhookServer() {
	if err := http.ListenAndServe(b.Config.Server.Host+b.Config.Server.Port, nil); err != nil {
		b.Logger.Fatal("failed start http server", zap.String("error", err.Error()))
	}
}

func (b *Bot) SendMessage(msg tgbotapi.Chattable) {
	_, err := b.API.Request(msg)
	if err != nil {
		b.Logger.Error("failed send message to bot", zap.String("error", err.Error()))
	}
}
