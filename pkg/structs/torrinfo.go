package structs

type TorrInfo struct {
	Data struct {
		Category  string `json:"category"`
		Uploader  string `json:"uploader"`
		Seeders   int    `json:"seeders"`
		Leechers  int    `json:"leechers"`
		Size      string `json:"size"`
		Downloads int    `json:"downloads"`
		Title     string `json:"title"`
		Link      string `json:"link"`
		Torrent   string `json:"torrent"`
		Magnet    string `json:"magnet"`
	} `json:"data"`
}
