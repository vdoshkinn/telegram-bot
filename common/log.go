package common

import (
	"log"

	"github.com/go-telegram/bot/models"
)

func DefaultLogging(message *models.Message, error error) {
	if error != nil {
		if message != nil && len(message.Text) > 0 {
			log.Println(message.Text)
		}

		log.Println(error.Error())
	}

	if message != nil && len(message.Text) > 0 {
		log.Println(message.Text)
	}
}
