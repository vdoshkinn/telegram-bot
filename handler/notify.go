package handler

import (
	"context"
	"telegram-bot/common"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func NotifyHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	common.SendMessageToAdminChannel(ctx, b, update, "Гость выехал")
	common.SendMessage(ctx, b, update, "Спасибо за выбор нашей квартиры для проживания. Будем рады видеть Вас снова!")
	//TODO сделать опросник и кидать ссылку на отзывы, если оценка 5. Если не 5 - узнать фидбек.
}
