package repository

import (
	"context"

	"github.com/eikoshelev/go-telegram-bot-template/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type WelcomeRepo struct {
	db pgxpool.Pool
}

const createUserQuery = `
INSERT INTO public.user_profile (tg_user_id, tg_chat_id, created_at)
VALUES ($1, $2, $3);
`

func (wr *WelcomeRepo) CreateProfile(userProfile *model.UserProfile) error {
	_, err := wr.db.Exec(
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
