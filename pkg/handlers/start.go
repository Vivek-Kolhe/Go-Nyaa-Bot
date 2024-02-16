package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var startMessage = bot.EscapeMarkdown("Hello!\n") +
	"I can fetch torrents from [Nyaa](https://nyaa.si/) and [Sukebei](https://sukebei.nyaa.si/)" +
	bot.EscapeMarkdown(".\nType /help for more info.")

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      startMessage,
		ParseMode: models.ParseModeMarkdown,
		LinkPreviewOptions: &models.LinkPreviewOptions{
			IsDisabled: bot.True(),
		},
	})
}
