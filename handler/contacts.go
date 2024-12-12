package handler

import (
	"context"
	"telegram-bot/common"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func MyContactsHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	text := "Если у Вас возник вопрос, смело звоните нам. Будем рады Вам помочь\n" +
		"Ольга: 89198340757\n" +
		"Диана: 89172031051"
	common.SendMessage(ctx, b, update, text)
}
