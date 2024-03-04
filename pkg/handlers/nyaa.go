package handlers

import (
	"context"
	"strings"

	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var cats = []string{"Anime", "Manga", "Audio", "Pics", "Live Action", "Software"}

var subCatMap = map[string][]string{
	"anime":       []string{"AMV", "Eng", "Non-Eng", "Raw"},
	"manga":       []string{"Eng", "Non-Eng", "Raw"},
	"audio":       []string{"Lossy", "Lossless"},
	"pics":        []string{"Photos", "Graphics"},
	"live action": []string{"Promo", "Eng", "Non-Eng", "Raw"},
	"software":    []string{"Applications", "Games"},
}

func NyaaHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	msgSlice := strings.SplitN(update.Message.Text, " ", 2)
	if len(msgSlice) < 2 {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Missing search query!",
		})
		return
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Please choose one of the following categories: ",
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: utils.GenerateCatBtns(cats, "nyaa", msgSlice[1]),
		},
	})
}

func NyaaCatCallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})

	callbackSlice := strings.Split(update.CallbackQuery.Data, " #$ ")

	if len(callbackSlice) == 3 {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Chat.ID,
			// MessageID: update.CallbackQuery.Message.MessageID,
			Text: "Choose one of the following sub-categories: ",
			ReplyMarkup: &models.InlineKeyboardMarkup{
				InlineKeyboard: utils.GenerateSubCatBtns(subCatMap, callbackSlice),
			},
		})
		return
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Chat.ID,
		// MessageID: update.CallbackQuery.Message.MessageID,
		Text: update.CallbackQuery.Data,
	})
}
