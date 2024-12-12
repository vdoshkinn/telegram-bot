package common

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func SendMessage(ctx context.Context, b *bot.Bot, update *models.Update, text string) {
	message, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      text,
		ParseMode: models.ParseModeMarkdownV1,
	})
	DefaultLogging(message, err)
}

func SendMessageToChannel(ctx context.Context, b *bot.Bot, channelId int, text string) {
	message, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    channelId,
		Text:      text,
		ParseMode: models.ParseModeMarkdownV1,
	})
	DefaultLogging(message, err)
}

func SendMessageWithReply(ctx context.Context, b *bot.Bot, update *models.Update, text string, kb *models.InlineKeyboardMarkup) {
	message, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        text,
		ParseMode:   models.ParseModeMarkdownV1,
		ReplyMarkup: kb,
	})
	DefaultLogging(message, err)
}

func SendMessageToAdminChannel(ctx context.Context, b *bot.Bot, update *models.Update, text string) {
	var chatId int64
	if update.Message != nil {
		chatId = update.Message.Chat.ID
	} else if update.CallbackQuery != nil {
		chatId = update.CallbackQuery.Message.Message.Chat.ID
	}

	text = fmt.Sprintf("%s:%d\n%s", ChatIdTeg, chatId, text)
	message, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    AdminChatId,
		Text:      text,
		ParseMode: models.ParseModeMarkdownV1,
	})
	DefaultLogging(message, err)
}

func ResendPhoto(ctx context.Context, b *bot.Bot, chatId int, update *models.Update) {
	if len(update.Message.Photo) > 0 {
		file, _ := b.GetFile(ctx, &bot.GetFileParams{
			FileID: update.Message.Photo[0].FileID,
		},
		)
		b.SendPhoto(ctx, &bot.SendPhotoParams{
			ChatID: chatId,
			Photo: &models.InputFileString{
				Data: file.FileID,
			},
		})
	}
}

func ResendDocument(ctx context.Context, b *bot.Bot, chatId int, update *models.Update) {
	if update.Message.Document != nil {
		file, _ := b.GetFile(ctx, &bot.GetFileParams{
			FileID: update.Message.Document.FileID,
		},
		)
		b.SendDocument(ctx, &bot.SendDocumentParams{
			ChatID: chatId,
			Document: &models.InputFileString{
				Data: file.FileID,
			},
		})
	}
}

func SetReaction(ctx context.Context, b *bot.Bot, update *models.Update) {
	reaction, err := b.SetMessageReaction(ctx, &bot.SetMessageReactionParams{
		ChatID:    update.Message.Chat.ID,
		MessageID: update.Message.ID,
		Reaction: []models.ReactionType{
			{
				Type: models.ReactionTypeTypeEmoji,
				ReactionTypeEmoji: &models.ReactionTypeEmoji{
					Type:  models.ReactionTypeTypeEmoji,
					Emoji: "👍",
				},
			},
		},
	})
	if err != nil {
		log.Println(err)
	}
	log.Println("Reaction was set:" + strconv.FormatBool(reaction))
}
