package constants

import "github.com/go-telegram/bot"

// This file contains all the text messages used. Only reason for having this in separate file is these can't be declared as consts and looks really awful.
// I found this workaround as passing the complete as string doesn't work, I don't know if that was intentional or not.
// Gotta find a cleaner way to do this.

// Messages used with start command.
var StartMessage = bot.EscapeMarkdown("Hello!\n") +
	"I can fetch torrents from [Nyaa](https://nyaa.si/) and [Sukebei](https://sukebei.nyaa.si/)" +
	bot.EscapeMarkdown(".\nType /help for more info.")

// Messages used with help command.
var HelpMessage = "*Note:* " +
	bot.EscapeMarkdown("The bot will fetch some of the recent torrents, so be specific with search query.\n\n") +
	"*Available commands:*" +
	bot.EscapeMarkdown("\n- /start - To check whether the bot is alive.\n- /help - To display this message.\n- /magnet - To get torrent info from ") +
	"*Nyaa* and *Sukebei* " +
	bot.EscapeMarkdown("using ID.")

var NyaaHelpMessage = "*For searching on [Nyaa](https://nyaa.si)*:" +
	bot.EscapeMarkdown("\n\n- /anime - Anime torrents.\n- /manga - Manga torrents.\n- /audio - Audio/Soundtrack torrents.\n- /live_action - Live Action shows/movies torrents.\n- /pictures - Picturebook/Graphics torrents.\n- /software - Games/Applications torrents.")

var SukebeiHelpMessage = "*For searching on [Sukebei](https://sukebei.nyaa.si)*:" +
	bot.EscapeMarkdown("\n\n- /art - Art torrents (anime, manga, doujinshi etc.).\n- /real - Non-Anime torrents.")
