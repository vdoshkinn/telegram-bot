package handler

import (
	"bufio"
	"context"
	"errors"
	"log"
	"os"
	"regexp"
	"telegram-bot/common"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

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
	file, err := os.Open("pin.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func SavePin(pin string) error {
	re := regexp.MustCompile("[0-9]+")
	pinDigits := re.FindAllString(pin, -1)
	if len(pinDigits) < 1 && len(pinDigits[0]) != 4 {
		return errors.New("invalid length for pin")
	}
	err := os.WriteFile("pin.txt", []byte(pinDigits[0]), 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
