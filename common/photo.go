package common

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func SendPhoto(ctx context.Context, b *bot.Bot, update *models.Update, photoName string) {
	fileContent, _ := os.ReadFile(fmt.Sprintf("photo/%s", photoName))
	params := &bot.SendPhotoParams{
		ChatID: update.Message.Chat.ID,
		Photo:  &models.InputFileUpload{Filename: photoName, Data: bytes.NewReader(fileContent)},
	}
	message, err := b.SendPhoto(ctx, params)
	DefaultLogging(message, err)
}
