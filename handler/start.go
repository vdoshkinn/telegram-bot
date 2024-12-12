package handler

import (
	"context"
	"telegram-bot/common"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	text := "Рады Вас приветствовать в чат-боте Апартаментов 'Счастливый слон'. \n" +
		"Для связи с администратором, можете просто написать сообщение в чат с ботом, администратор его увидит и сможет ответить.\n" +
		"Для наиболее частых запросов в боте есть команды, которые позволят вам получить моментальный ответ.\n" +
		"Информация о том, как добраться до квартиры находится в команде '\\info'"
	common.SendMessage(ctx, b, update, text)
}
