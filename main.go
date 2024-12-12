package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"telegram-bot/handler"

	"github.com/go-telegram/bot"
)

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler.DefaultHandler),
		bot.WithCallbackQueryDataHandler("bank", bot.MatchTypePrefix, handler.CallbackHandler),
	}

	b, err := bot.New("7859184648:AAG2M5oKfrY7pALwxK7D_88lCCdAAPphQ2c", opts...)
	if err != nil {
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/info", bot.MatchTypePrefix, handler.MyInfoHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypePrefix, handler.StartHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/wifi", bot.MatchTypePrefix, handler.MyWifiHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/contacts", bot.MatchTypePrefix, handler.MyContactsHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/rules", bot.MatchTypePrefix, handler.MyRulesHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/notify", bot.MatchTypePrefix, handler.NotifyHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/document", bot.MatchTypePrefix, handler.RequestDocumentHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/checkout", bot.MatchTypePrefix, handler.CheckoutHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/features", bot.MatchTypePrefix, handler.FeaturesHandler)

	log.Println("Bot is now running.  Press CTRL-C to exit.")
	b.Start(ctx)

}
