package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"telegram-bot/common"
	"telegram-bot/cron"
	"telegram-bot/handler"

	"github.com/go-telegram/bot"
)

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var err error
	defer func() {
		if panicErr := recover(); panicErr != nil {
			log.Println(fmt.Sprintf("Panic occured %v", panicErr))
		}

		if err != nil {
			log.Println(fmt.Sprintf("Error occured %v", err))
		}
	}()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler.DefaultHandler),
		bot.WithCallbackQueryDataHandler("bank", bot.MatchTypePrefix, handler.CallbackHandler),
	}

	var b *bot.Bot
	b, err = bot.New("7859184648:AAG2M5oKfrY7pALwxK7D_88lCCdAAPphQ2c", opts...)
	if err != nil {
		panic(err)
	}

	service := common.NewPermissionsService(0)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/info", bot.MatchTypePrefix, handler.MyInfoHandler, service.CheckUserIsAcceptable)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypePrefix, handler.StartHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/wifi", bot.MatchTypePrefix, handler.MyWifiHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/contacts", bot.MatchTypePrefix, handler.MyContactsHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/rules", bot.MatchTypePrefix, handler.MyRulesHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/notify", bot.MatchTypePrefix, handler.NotifyHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/document", bot.MatchTypePrefix, handler.RequestDocumentHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/checkout", bot.MatchTypePrefix, handler.CheckoutHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/features", bot.MatchTypePrefix, handler.FeaturesHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/pin", bot.MatchTypePrefix, handler.MyPinHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/add", bot.MatchTypePrefix, common.EmptyHandler, service.AddUser)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/remove", bot.MatchTypePrefix, common.EmptyHandler, service.RemoveAllUsers)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/admin", bot.MatchTypePrefix, handler.AdminHandler)

	log.Println("Bot is now running.  Press CTRL-C to exit.")

	go cron.RunCron(ctx, b)
	b.Start(ctx)

}
