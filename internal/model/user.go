package model

import (
	"time"
)

type (
	TelegramUserID int64
	TelegramChatID int64
)

type UserProfile struct {
	TelegramUserID TelegramUserID
	TelegramChatID TelegramChatID
	CreatedAt      time.Time
}
