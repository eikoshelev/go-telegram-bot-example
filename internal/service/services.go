package service

import (
	"github.com/eikoshelev/go-telegram-bot-template/internal/repository"
)

type Service struct {
	Start   *Start
	Profile *Profile
	// list all your services here
}

// function to initialize all services
func Init(repos *repository.Repository) *Service {
	return &Service{
		Start:   &Start{Start: repos.Start},
		Profile: &Profile{Profile: repos.Profile},
		// don't forget to inject them as a dependency
	}
}
