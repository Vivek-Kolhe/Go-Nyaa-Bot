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
	var url string
	callbackData := strings.Split(update.CallbackQuery.Data, " #$ ")
	site, torrID := callbackData[1], callbackData[2]

	switch site {
	case "nyaa":
		url = fmt.Sprintf("%s%s", constants.NyaaMagnet, torrID)
	case "sukebei":
		url = fmt.Sprintf("%s%s", constants.SukebeiMagnet, torrID)
	}

	bytes, err := utils.MakeRequest(url)
	if err != nil {
		log.Panic(err.Error())
	}

	var data *structs.TorrInfo
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Panic(err.Error())
	}

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    update.CallbackQuery.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.MessageID,
		Text:      utils.GenerateTorrInfoMsg(data),
		ParseMode: models.ParseModeMarkdown,
	})
}
