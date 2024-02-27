package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/handlers"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("D:\\Code\\Go Workspace\\src\\github.com\\Vivek-Kolhe\\gonyaa-bot\\.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handlers.DefaultHandler),
		bot.WithCallbackQueryDataHandler("help", bot.MatchTypePrefix, handlers.HelpCallbackHandler),
		bot.WithCallbackQueryDataHandler("magnet", bot.MatchTypePrefix, handlers.MagnetCallbackHandler),
	}

	b, err := bot.New(os.Getenv("BOT_TOKEN"), opts...)
	if err != nil {
		log.Panic(err.Error())
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers.StartHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, handlers.HelpHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/magnet", bot.MatchTypePrefix, handlers.MagnetHandler)
	b.Start(ctx)
}
