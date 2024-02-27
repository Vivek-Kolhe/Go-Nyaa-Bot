package handlers

import (
	"context"
	"strings"

	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/constants"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HelpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        constants.HelpMessage,
		ReplyMarkup: constants.HelpButtons,
		ParseMode:   models.ParseModeMarkdown,
	})
}

func HelpCallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	callbackData := strings.Split(update.CallbackQuery.Data, " #$ ")[1]
	switch callbackData {
	case "nyaa":
		handleHelpCallback(ctx, b, update, constants.NyaaHelpMessage, constants.BackButton)
	case "sukebei":
		handleHelpCallback(ctx, b, update, constants.SukebeiHelpMessage, constants.BackButton)
	case "back":
		handleHelpCallback(ctx, b, update, constants.HelpMessage, constants.HelpButtons)
	}
}

func handleHelpCallback(ctx context.Context, b *bot.Bot, update *models.Update, msg string, replyMarkUp *models.InlineKeyboardMarkup) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      update.CallbackQuery.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.MessageID,
		Text:        msg,
		ParseMode:   models.ParseModeMarkdown,
		ReplyMarkup: replyMarkUp,
		LinkPreviewOptions: &models.LinkPreviewOptions{
			IsDisabled: bot.True(),
		},
	})
}
