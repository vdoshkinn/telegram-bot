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

func SavePin(command string) error {
	digits, err := parseDigits(command)
	if err != nil {
		return err
	}

	return common.SaveStringToFile(path, digits)
}

func parseDigits(command string) (string, error) {
	res := regexp.MustCompile("[0-9]+")
	pinDigits := res.FindAllString(command, -1)
	if len(pinDigits) < 1 && len(pinDigits[0]) != 4 {
		return "", errors.New("invalid length for command")
	}
	return pinDigits[0], nil
}
