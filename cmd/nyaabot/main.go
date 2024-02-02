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

	token := os.Getenv("BOT_TOKEN")
	b, err := bot.New(token)
	if err != nil {
		panic(err.Error())
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers.HelloHandler)
	b.Start(ctx)
}
