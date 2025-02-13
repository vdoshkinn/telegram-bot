package cron

import (
	"context"
	"log"
	"telegram-bot/common"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/robfig/cron/v3"
)

func RunCron(ctx context.Context, b *bot.Bot) {
	// Устанавливаем часовой пояс на московское время
	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatalf("Ошибка загрузки локации: %v", err)
	}

	c := cron.New(cron.WithLocation(location))

	// Добавляем задачу, которая будет выполняться каждый день в 10:00 по московскому времени
	_, err = c.AddFunc("0 11 * * *", func() {
		log.Println("check")
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    common.AdminChatId,
			Text:      "check",
			ParseMode: models.ParseModeMarkdownV1,
		})
	})

	if err != nil {
		log.Fatalf("Ошибка добавления задачи: %v", err)
	}

	// Запускаем cron-расписатель
	c.Start()
	log.Println("Крон запущен")

	select {}
}
