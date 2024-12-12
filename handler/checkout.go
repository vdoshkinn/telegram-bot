package handler

import (
	"context"
	"telegram-bot/common"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var idToBankNameMap = map[string]string{
	"bank_1": "ВТБ",
	"bank_2": "Сбер",
	"bank_3": "Райфайзен",
	"bank_4": "Тинькофф",
	"bank_5": "Другой банк",
}

const otherBank = "bank_5"

func CheckoutHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	text := "При выезде (напомню, что он осуществляется до 12.00, если будете выезжать раньше - сообщите пожалуйста):\n" +
		"1. Оставьте ключ в сейфе возле двери и сменить комбинацию цифр\n" +
		"2. Вызвать команду бота \\notify\n" +
		"3. Сообщить банк для возврата залога\n"

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: idToBankNameMap["bank_1"], CallbackData: "bank_1"},
				{Text: idToBankNameMap["bank_2"], CallbackData: "bank_2"},
			}, {
				{Text: idToBankNameMap["bank_3"], CallbackData: "bank_3"},
				{Text: idToBankNameMap["bank_4"], CallbackData: "bank_4"},
			},
			{
				{Text: idToBankNameMap["bank_5"], CallbackData: "bank_5"},
			},
		},
	}
	common.SendMessageWithReply(ctx, b, update, text, kb)
}

func CallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// answering callback query first to let Telegram know that we received the callback query,
	// and we're handling it. Otherwise, Telegram might retry sending the update repetitively
	// as it thinks the callback query doesn't reach to our application. learn more by
	// reading the footnote of the https://core.telegram.org/bots/api#callbackquery type.
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	if update.CallbackQuery.Data == otherBank {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Text:   "Просто напишите название банка",
		})
		common.SendMessageToAdminChannel(ctx, b, update, "Пользователь выбрал другой банк")
	} else {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Text:   "Мы вернем вам залог на банк: " + idToBankNameMap[update.CallbackQuery.Data],
		})
		common.SendMessageToAdminChannel(ctx, b, update, "Залог возвращать на "+idToBankNameMap[update.CallbackQuery.Data])
	}
}
