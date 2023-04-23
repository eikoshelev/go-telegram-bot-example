package repository

import (
	"context"

	"github.com/eikoshelev/go-telegram-bot-example/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Profile struct {
	db pgxpool.Pool
}

// example request to create a user
const createUserQuery = `
INSERT INTO public.user_profile (tg_user_id, tg_chat_id, created_at)
VALUES ($1, $2, $3);
`

// example function to create a user
func (p *Profile) Create(userProfile *model.UserProfile) error {
	_, err := p.db.Exec(
		context.Background(),
		createUserQuery,
		userProfile.TelegramUserID,
		userProfile.TelegramChatID,
	)
	if err != nil {
		return err
	}
	return nil
}
