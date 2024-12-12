package handler

import (
	"context"
	"telegram-bot/common"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func RequestDocumentHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	common.SendMessageToAdminChannel(ctx, b, update, "Гостю нужна справка")
	common.SendMessage(ctx, b, update, "Мы сделаем справку в течении суток и свяжемся с Вами")
}
