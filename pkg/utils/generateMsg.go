package utils

import (
	"fmt"
	"strings"

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

// Function for generating message for torrent listing using Torrents struct. Used in search handler.
func GenerateTorrListMsg(data structs.Torrents) []string {
	msgs := make([]string, 0)
	for i := 0; i < min(data.Count, 10); i++ {
		link := strings.Split(data.Data[i].Link, "/")
		id := link[len(link)-1]
		msg := "*" + bot.EscapeMarkdown(data.Data[i].Title) + "*" +
			"\n*ID: *" + "`" + bot.EscapeMarkdown(id) + "`" +
			"\n*Size: *" + "`" + bot.EscapeMarkdown(data.Data[i].Size) + "`" +
			"\n*Seeders/Leechers: *" + "`" + bot.EscapeMarkdown(fmt.Sprintf("%d/%d", data.Data[i].Seeders, data.Data[i].Leechers)) + "`"

		msgs = append(msgs, msg)
	}
	return msgs
}
