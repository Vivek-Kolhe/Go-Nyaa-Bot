package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/constants"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func PingHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	start := time.Now().UnixMilli()
	msg, _ := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Pong!",
	})
	end := time.Now().UnixMilli()
	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    update.Message.Chat.ID,
		MessageID: msg.ID,
		Text:      fmt.Sprintf("%s `%d ms`", constants.PingMessage, end-start),
		ParseMode: models.ParseModeMarkdown,
	})
}
