package constants

import "github.com/go-telegram/bot/models"

// This file contains all the InlineButtons used.

// InlineButtons used with help command.
var HelpButtons = &models.InlineKeyboardMarkup{
	InlineKeyboard: [][]models.InlineKeyboardButton{
		{
			{Text: "Nyaa", CallbackData: "nyaahelp"},
			{Text: "Sukebei", CallbackData: "sukebeihelp"},
		},
	},
}

var BackButton = &models.InlineKeyboardMarkup{
	InlineKeyboard: [][]models.InlineKeyboardButton{
		{
			{Text: "Back", CallbackData: "backhelp"},
		},
	},
}
