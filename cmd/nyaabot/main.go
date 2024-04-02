package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/handlers"
	"github.com/go-telegram/bot"
)

// func init() {
// 	err := godotenv.Load("..\\..\\.env")
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	http.HandleFunc("/", Hello)

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
	b.RegisterHandler(bot.HandlerTypeMessageText, "/ping", bot.MatchTypeExact, handlers.PingHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/magnet", bot.MatchTypePrefix, handlers.MagnetHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/nyaa", bot.MatchTypePrefix, handlers.SearchHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/sukebei", bot.MatchTypePrefix, handlers.SearchHandler)

	go func() {
		b.Start(ctx)
	}()

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err.Error())
	}
}
