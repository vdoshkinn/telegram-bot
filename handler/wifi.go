package handler

import (
	"context"
	"telegram-bot/common"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func MyWifiHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	text := "Название сети: HappyElephant\n" +
		"Пароль: Appartements137"
	common.SendMessage(ctx, b, update, text)

	common.SendPhoto(ctx, b, update, "qr_wifi.jpg")
}
