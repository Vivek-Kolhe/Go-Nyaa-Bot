package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/constants"
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
	} else {
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
}
