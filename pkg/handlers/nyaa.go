package handlers

import (
	"context"
	"strings"

	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var cats = []string{"Anime", "Manga", "Audio", "Pics", "LA", "Software"}

// var subCatMap = map[string][]string{
// 	"Anime":    []string{"AMV", "Eng", "Non-Eng", "Raw"},
// 	"Manga":    []string{"Eng", "Non-Eng", "Raw"},
// 	"Audio":    []string{"Lossy", "Lossless"},
// 	"Pics":     []string{"Photos", "Graphics"},
// 	"LA":       []string{"Promo", "Eng", "Non-Eng", "Raw"},
// 	"Software": []string{"Applications", "Games"},
// }

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

func NyaaCallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Chat.ID,
		Text:   update.CallbackQuery.Data,
	})
}
