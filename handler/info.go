package handler

import (
	"context"
	"telegram-bot/common"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func MyInfoHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	text := "Квартира находится по адресу г. Мурино, проспект Авиаторов Балтики, 11/1, кв 137.\n" +
		"[Яндекс Карты](https://yandex.ru/maps/118936/murino/house/prospekt_aviatorov_baltiki_11_1/Z0kYcgNgSEAOQFhqfXx0dH1iZQ==) \n" +
		"[2Гис](https://2gis.ru/spb/geo/5348660212853014/30.439654%2C60.054387/entranceId/70030076187672376?m=30.439938%2C60.054317%2F19.03) \n" +
		"[Маршрут от метро](https://yandex.ru/maps/-/CHAtiN9y)"
	common.SendMessage(ctx, b, update, text)
}
