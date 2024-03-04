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

func GenerateSubCatBtns(subCatMap map[string][]string, callbackSlice []string) [][]models.InlineKeyboardButton {
	inlineKb := make([][]models.InlineKeyboardButton, 0)
	site, cat, query := callbackSlice[0], callbackSlice[1], callbackSlice[2]

	curr := subCatMap[cat]
	for i := 0; i < len(curr); i += 2 {
		temp := []models.InlineKeyboardButton{}
		for j := i; j < min(i+2, len(curr)); j++ {
			btn := models.InlineKeyboardButton{
				Text:         curr[j],
				CallbackData: fmt.Sprintf("%s #$ %s #$ %s #$ %s", site, cat, strings.ToLower(curr[j]), query),
			}
			temp = append(temp, btn)
		}
		inlineKb = append(inlineKb, temp)
	}
	return inlineKb
}
