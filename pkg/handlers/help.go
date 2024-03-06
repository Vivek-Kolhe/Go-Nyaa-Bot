package handlers

import (
	"context"

	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/constants"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HelpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      constants.HelpMessage,
		ParseMode: models.ParseModeMarkdown,
		LinkPreviewOptions: &models.LinkPreviewOptions{
			IsDisabled: bot.True(),
		},
	})
}
