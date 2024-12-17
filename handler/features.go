package handler

import (
	"context"
	"telegram-bot/common"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func FeaturesHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	text := "Второй ключ находится в консоли на входе.\n" +
		"Сушилку и гладильную доска Вы сможете найти на балконе.\n" +
		"Также на балконе находится *закрытый шкаф* 🔒 для хозяйственных нужд.\n" +
		"Фен находится в тумбе в ванной комнате.\n" +
		"Утюг находится в шкафу в коридоре.\n" +
		"В том же шкафу находится пленка, которой можно воспользоваться для упаковки чемоданов."

	common.SendMessage(ctx, b, update, text)
}
