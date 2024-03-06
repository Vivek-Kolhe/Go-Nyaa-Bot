package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/constants"
	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/structs"
	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var (
	cats = map[string][]string{
		"nyaa":    {"Anime", "Manga", "Audio", "Pictures", "Live Action", "Software"},
		"sukebei": {"Art", "Real"},
	}

	subCats = map[string]map[string][]string{
		"nyaa": {
			"anime":       {"AMV", "Eng", "Non-Eng", "Raw"},
			"manga":       {"Eng", "Non-Eng", "Raw"},
			"audio":       {"Lossy", "Lossless"},
			"pictures":    {"Photos", "Graphics"},
			"live action": {"Promo", "Eng", "Non-Eng", "Raw"},
			"software":    {"Applications", "Games"},
		},
		"sukebei": {
			"art":  {"Anime", "Doujinshi", "Games", "Manga", "Pictures"},
			"real": {"Photos", "Videos"},
		},
	}
)

func SearchHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	msgSlice := strings.SplitN(update.Message.Text, " ", 2)
	if len(msgSlice) < 2 {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Missing search query!",
		})
		return
	}

	site := msgSlice[0][1:]
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Please choose one of the following categories: ",
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: utils.GenerateCatBtns(cats[site], site, msgSlice[1]),
		},
	})
}

func SearchCallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
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
				InlineKeyboard: utils.GenerateSubCatBtns(subCats[callbackSlice[0]], callbackSlice),
			},
		})
		return
	}

	params := url.Values{}
	params.Set("category", strings.ReplaceAll(callbackSlice[1], " ", "_"))
	params.Set("sub_category", callbackSlice[2])
	params.Set("q", callbackSlice[3])

	url := fmt.Sprintf("%s?%s", constants.SearchEndpoint[callbackSlice[0]], params.Encode())
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

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    update.CallbackQuery.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.MessageID,
		Text:      strings.Join(utils.GenerateTorrListMsg(data), "\n\n"),
		ParseMode: models.ParseModeMarkdown,
	})
}
