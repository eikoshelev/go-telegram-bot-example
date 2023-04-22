package flow

import (
	"github.com/eikoshelev/go-telegram-bot-template/internal/messages/welcome"
	"github.com/eikoshelev/go-telegram-bot-template/internal/model"
	"github.com/eikoshelev/go-telegram-bot-template/internal/service"
)

func Init(svc *service.Service) model.Flow {
	return model.Flow{
		model.StartCommand: model.Usecase{
			"welcome": model.Chain{
				0: model.Action{
					Handler: svc.Welcome.WelcomeHandler,
					Message: welcome.Welcome(),
				},
				1: model.Action{
					Handler: svc.Welcome.CreateProfile,
					Message: welcome.CreateProfile(),
				},
			},
		},
		/*
			model.<CommandName>: model.Usecase{
				"<Some case>": model.Chain{
					0: model.Action{
						Handler: <Service handler>,
						Message: <Some message>,
					},
					1: model.Action{
						Handler: <Service handler>,
						Message: <Some message>,
					},
					2: model.Action{
						Handler: <Service handler>,
						Message: <Some message>,
					},
				},
			},
		*/
	}
}
