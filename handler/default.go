package handler

import (
	"context"
	"strconv"
	"strings"
	"telegram-bot/common"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message.Chat.ID == common.AdminChatId {
		replyFromAdminChat(ctx, b, update)
		return
	}

	common.SetReaction(ctx, b, update.Message.ID, update.Message.Chat.ID)
	common.SendMessageToAdminChannel(ctx, b, update, update.Message.Text)
	common.ResendPhoto(ctx, b, common.AdminChatId, update)
	common.ResendDocument(ctx, b, common.AdminChatId, update)
}

func replyFromAdminChat(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message.ReplyToMessage != nil && strings.Contains(update.Message.ReplyToMessage.Text, common.ChatIdTeg) {
		firstLine := strings.Split(update.Message.ReplyToMessage.Text, "\n")[0]
		chatID := strings.Split(firstLine, ":")[1]
		chatIDint, _ := strconv.Atoi(chatID)
		_, err := common.SendMessageToChannel(ctx, b, chatIDint, update.Message.Text)
		if err == nil {
			common.SetReaction(ctx, b, update.Message.ID, update.Message.Chat.ID)
		}
		common.ResendPhoto(ctx, b, chatIDint, update)
		common.ResendDocument(ctx, b, chatIDint, update)
	}
}
