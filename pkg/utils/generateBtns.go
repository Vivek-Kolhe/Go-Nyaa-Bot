package utils

import (
	"fmt"
	"strings"

	"github.com/go-telegram/bot/models"
)

func GenerateCatBtns(catMap []string, site, callbackData string) [][]models.InlineKeyboardButton {
	inlineKb := make([][]models.InlineKeyboardButton, 0)

	for i := 0; i < len(catMap); i += 2 {
		temp := []models.InlineKeyboardButton{}
		for j := i; j < min(i+2, len(catMap)); j++ {
			btn := models.InlineKeyboardButton{
				Text:         catMap[j],
				CallbackData: fmt.Sprintf("%s #$ %s #$ %s", site, strings.ToLower(catMap[j]), callbackData),
			}
			temp = append(temp, btn)
		}
		inlineKb = append(inlineKb, temp)
	}
	return inlineKb
}
