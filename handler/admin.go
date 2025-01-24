package handler

import (
	"context"
	"telegram-bot/common"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const adminPath = "admin_commands.txt"

func AdminHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message.Chat.ID != common.AdminChatId {
		return
	}
	common.SendTextToAdminChannel(ctx, b, common.ReadStringFromFile(adminPath))
}
