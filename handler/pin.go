package handler

import (
	"context"
	"errors"
	"regexp"
	"telegram-bot/common"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const path = "pin.txt"

func MyPinHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message.Chat.ID != common.AdminChatId {
		return
	}
	err := SavePin(update.Message.Text)
	if err == nil {
		common.SetReaction(ctx, b, update.Message.ID, update.Message.Chat.ID)
	}
}

func GetPin() string {
	return common.ReadStringFromFile(path)
}

func SavePin(pin string) error {
	re := regexp.MustCompile("[0-9]+")
	pinDigits := re.FindAllString(pin, -1)
	if len(pinDigits) < 1 && len(pinDigits[0]) != 4 {
		return errors.New("invalid length for pin")
	}
	return common.SaveStringToFile(path, pin)
}
