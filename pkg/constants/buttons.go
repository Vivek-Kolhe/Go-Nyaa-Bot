package constants

import "github.com/go-telegram/bot/models"

// This file contains all the InlineButtons used. Only reason for having this in separate file is these can't be declared as consts.

// InlineButtons used with help command.
var HelpButtons = &models.InlineKeyboardMarkup{
	InlineKeyboard: [][]models.InlineKeyboardButton{
		{
			{Text: "Nyaa", CallbackData: "help #$ nyaa"},
			{Text: "Sukebei", CallbackData: "help #$ sukebei"},
		},
	},
}

var BackButton = &models.InlineKeyboardMarkup{
	InlineKeyboard: [][]models.InlineKeyboardButton{
		{
			{Text: "Back", CallbackData: "help #$ back"},
		},
	},
}
