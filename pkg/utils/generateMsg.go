package utils

import (
	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/structs"
	"github.com/go-telegram/bot"
)

func GenerateTorrInfoMsg(data *structs.TorrInfo) string {
	msg := "*" + bot.EscapeMarkdown(data.Data.Title) + "*"
	return msg
}
