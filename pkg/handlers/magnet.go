package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/constants"
	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/structs"
	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func MagnetHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	msgSlice := strings.Split(update.Message.Text, " ")
	if len(msgSlice) < 2 {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   constants.NoIDMessage,
		})
		return
	}

	buttons := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Nyaa", CallbackData: fmt.Sprintf("magnet #$ nyaa #$ %s", msgSlice[1])},
				{Text: "Sukebei", CallbackData: fmt.Sprintf("magnet #$ sukebei #$ %s", msgSlice[1])},
			},
		},
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        constants.MagnetMessage,
		ReplyMarkup: buttons,
	})
}

func MagnetCallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})

	callbackData := strings.Split(update.CallbackQuery.Data, " #$ ")
	site, torrID := callbackData[1], callbackData[2]
	url := fmt.Sprintf("%s%s", constants.MagnetEndpoint[site], torrID)

	bytes, statusCode, err := utils.MakeRequest(url)
	if statusCode != 200 || err != nil {
		if statusCode == 404 || statusCode == 422 {
			b.EditMessageText(ctx, &bot.EditMessageTextParams{
				ChatID:    update.CallbackQuery.Message.Chat.ID,
				MessageID: update.CallbackQuery.Message.MessageID,
				Text:      constants.InvalidIDMessage,
			})
			return
		}
		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.MessageID,
			Text:      constants.SomethingWentWrong,
		})
		return
	}

	var data structs.TorrInfo
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Panic(err.Error())
	}

	linkBtn := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "View Torrent?", URL: data.Data.Link},
			},
		},
	}

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      update.CallbackQuery.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.MessageID,
		Text:        utils.GenerateTorrInfoMsg(data),
		ParseMode:   models.ParseModeMarkdown,
		ReplyMarkup: linkBtn,
	})
}
