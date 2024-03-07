package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/handlers"
	"github.com/go-telegram/bot"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handlers.DefaultHandler),
		bot.WithCallbackQueryDataHandler("magnet", bot.MatchTypePrefix, handlers.MagnetCallbackHandler),
		bot.WithCallbackQueryDataHandler("nyaa", bot.MatchTypePrefix, handlers.SearchCallbackHandler),
		bot.WithCallbackQueryDataHandler("sukebei", bot.MatchTypePrefix, handlers.SearchCallbackHandler),
	}

	b, err := bot.New(os.Getenv("BOT_TOKEN"), opts...)
	if err != nil {
		log.Panic(err.Error())
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers.StartHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, handlers.HelpHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/magnet", bot.MatchTypePrefix, handlers.MagnetHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/nyaa", bot.MatchTypePrefix, handlers.SearchHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/sukebei", bot.MatchTypePrefix, handlers.SearchHandler)
	b.Start(ctx)
}
