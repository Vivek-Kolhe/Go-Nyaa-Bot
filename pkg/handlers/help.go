package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var helpMessage = "*Note:* " +
	bot.EscapeMarkdown("The bot will fetch some of the recent torrents, so be specific with search query.\n\n") +
	"*Available commands:*" +
	bot.EscapeMarkdown("\n- /start - To check whether the bot is alive.\n- /help - To display this message.\n- /magnet - To get torrent info from ") +
	"*Nyaa* and *Sukebei* " +
	bot.EscapeMarkdown("using ID.")

var sukebeiHelpMessage = "*For searching on [Sukebei](https://sukebei.nyaa.si)*:" +
	bot.EscapeMarkdown("\n- /art - Art torrents (anime, manga, doujinshi etc.).\n- /real - Non-Anime torrents.")

var helpButtons = &models.InlineKeyboardMarkup{
	InlineKeyboard: [][]models.InlineKeyboardButton{
		{
			{Text: "Nyaa", CallbackData: "nyaahelp"},
			{Text: "Sukebei", CallbackData: "sukebeihelp"},
		},
	},
}

var backButton = &models.InlineKeyboardMarkup{
	InlineKeyboard: [][]models.InlineKeyboardButton{
		{
			{Text: "Back", CallbackData: "backhelp"},
		},
	},
}

func HelpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        helpMessage,
		ReplyMarkup: helpButtons,
		ParseMode:   models.ParseModeMarkdown,
	})
}

func SukebeiHelpCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      update.CallbackQuery.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.MessageID,
		Text:        sukebeiHelpMessage,
		ParseMode:   models.ParseModeMarkdown,
		ReplyMarkup: backButton,
		LinkPreviewOptions: &models.LinkPreviewOptions{
			IsDisabled: bot.True(),
		},
	})
}

func BackCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      update.CallbackQuery.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.MessageID,
		Text:        helpMessage,
		ReplyMarkup: helpButtons,
		ParseMode:   models.ParseModeMarkdown,
	})
}
