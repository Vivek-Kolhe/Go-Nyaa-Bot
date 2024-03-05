package handlers

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/constants"
	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/structs"
	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var cats = []string{"Anime", "Manga", "Audio", "Pics", "Live Action", "Software"}

var subCatMap = map[string][]string{
	"anime":       {"AMV", "Eng", "Non-Eng", "Raw"},
	"manga":       {"Eng", "Non-Eng", "Raw"},
	"audio":       {"Lossy", "Lossless"},
	"pics":        {"Photos", "Graphics"},
	"live action": {"Promo", "Eng", "Non-Eng", "Raw"},
	"software":    {"Applications", "Games"},
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
		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.MessageID,
			Text:      "Choose one of the following sub-categories: ",
			ReplyMarkup: &models.InlineKeyboardMarkup{
				InlineKeyboard: utils.GenerateSubCatBtns(subCatMap, callbackSlice),
			},
		})
		return
	}

	params := make(map[string]string)
	params["category"] = strings.ReplaceAll(callbackSlice[1], " ", "_")
	params["sub_category"] = callbackSlice[2]
	params["q"] = callbackSlice[3]

	url := utils.GenerateURL(constants.Nyaa, params)

	bytes, statusCode, err := utils.MakeRequest(url)
	if statusCode != 200 || err != nil {
		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.MessageID,
			Text:      url,
		})
		return
	}

	var data structs.Torrents
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Panic(err.Error())
	}

	if data.Count < 1 {
		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.MessageID,
			Text:      "No results found!",
		})
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Chat.ID,
		// MessageID: update.CallbackQuery.Message.MessageID,
		Text:      strings.Join(utils.GenerateTorrListMsg(data), "\n"),
		ParseMode: models.ParseModeMarkdown,
	})
}
