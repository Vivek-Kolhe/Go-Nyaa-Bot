package utils

import (
	"fmt"

	"github.com/Vivek-Kolhe/gonyaa-bot/pkg/structs"
	"github.com/go-telegram/bot"
)

// Function for generating message for torrent info using TorrInfo struct. Used in magnet handler.
func GenerateTorrInfoMsg(data structs.TorrInfo) string {
	msg := "*" + bot.EscapeMarkdown(data.Data.Title) + "*" +
		"\n*Seeders/Leechers: *" + "`" + bot.EscapeMarkdown(fmt.Sprintf("%d/%d", data.Data.Seeders, data.Data.Leechers)) + "`" +
		"\n*Size: *" + "`" + bot.EscapeMarkdown(data.Data.Size) + "`" +
		"\n*Uploader: *" + "`" + bot.EscapeMarkdown(data.Data.Uploader) + "`" +
		"\n*Torrent: *[" + bot.EscapeMarkdown("Click Here!") + "](" + bot.EscapeMarkdown(data.Data.Torrent) + ")" +
		"\n*Magnet: *" + "`" + bot.EscapeMarkdown(data.Data.Magnet) + "`"

	return msg
}

func GenerateTorrListMsg(data structs.Torrents) []string {
	result := make([]string, 0)
	for i := 0; i < min(data.Count, 20); i++ {
		temp := "*" + bot.EscapeMarkdown(data.Data[i].Title) + "*"
		result = append(result, temp)
	}
	return result
}
