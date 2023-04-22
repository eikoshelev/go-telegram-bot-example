package service

import (
	"github.com/eikoshelev/go-telegram-bot-template/internal/repository"
)

type Service struct {
	Welcome *WelcomeService
	// list all your services here
}

func Init(repos *repository.Repository) *Service {
	return &Service{
		Welcome: &WelcomeService{
			Welcome: repos.Welcome,
		},
		// don't forget to inject them as a dependency
	}
}
