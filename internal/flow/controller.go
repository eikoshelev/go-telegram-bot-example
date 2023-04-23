package flow

import (
	"github.com/eikoshelev/go-telegram-bot-template/internal/messages/profile"
	"github.com/eikoshelev/go-telegram-bot-template/internal/model"
	"github.com/eikoshelev/go-telegram-bot-template/internal/service"
)

func Init(svc *service.Service) model.Flow {
	return model.Flow{
		model.ProfileCommand: model.Usecase{
			"create": model.Chain{
				0: model.Action{
					Handler: svc.Profile.Create,
					Message: profile.Create(),
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
