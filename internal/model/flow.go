package model

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

/*
CommandFlow (starting a specific script (flow) for manipulating an object, is a command)
		  |
		  Usecase (actions that can be performed on an object)
				|
				Chain (algorithm, sequence of steps to implement an action and obtain some result)
					|
					Step (certain, specific step, action)

{
   "start":{
      "welcome":{
         "0":{
            "handler": HandlerFunc(),
            "message":"some text",
            "buttons":[
               {
                  "name":"button name 1",
                  "callback_data":{
                     "cmd_key":"start",
                     "case":"welcome",
                     "step":0,
                     "payload":"button 1 is pressed"
                  }
               }
            ]
         }
      }
   }
}
*/

const DefaltUsecase Case = "DefaltUsecase"

type (
	Flow    map[CommandKey]Usecase
	Usecase map[Case]Chain
	Case    string
	Chain   map[Step]Action
	Step    int
	Action  struct {
		Handler HandlerFunc
		Message
	}
	HandlerFunc func(*UpdateLocal) (tgbotapi.Chattable, error)
	Message     struct {
		Text    string
		Buttons []Button
	}
	Button struct {
		Name         string
		CallbackData CallbackData
	}
)

func (flow Flow) Handle(cd *CallbackData, updLocal *UpdateLocal) (tgbotapi.Chattable, error) {
	msg, err := flow[cd.CommandKey][cd.Case][cd.Step].Handler(updLocal)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (msg Message) BuildBotMessage(chatID int64) tgbotapi.MessageConfig {
	replyMessage := tgbotapi.NewMessage(chatID, msg.Text)
	var buttonRows [][]tgbotapi.InlineKeyboardButton
	for _, button := range msg.Buttons {
		buttonRows = append(buttonRows, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(button.Name, button.CallbackData.Encode()),
		),
		)
	}
	markup := tgbotapi.NewInlineKeyboardMarkup(
		buttonRows...,
	)
	replyMessage.ReplyMarkup = markup
	replyMessage.ParseMode = tgbotapi.ModeHTML
	return replyMessage
}

func (flow *Flow) ValidateCallbacksDataSize(logger *zap.Logger) {
	for _, usecase := range *flow {
		for _, chain := range usecase {
			for _, action := range chain {
				for _, button := range action.Buttons {
					// 64 bytes - telegram limit for callback_data: https://core.telegram.org/bots/api#inlinekeyboardbutton
					if len(button.CallbackData.Encode()) > 64 {
						logger.Fatal("size of callback_data exceeds 64 bytes", zap.String("callback_data", button.CallbackData.Encode()), zap.Int("bytes", len(button.CallbackData.Encode())))
					}
				}
			}
		}
	}
	logger.Info("callback_data dimensions are valid")
}
