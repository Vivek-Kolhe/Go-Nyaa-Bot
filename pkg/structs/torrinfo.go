package structs

type TorrInfo struct {
	Data struct {
		Uploader string `json:"uploader"`
		Seeders  int    `json:"seeders"`
		Leechers int    `json:"leechers"`
		Size     string `json:"size"`
		Title    string `json:"title"`
		Link     string `json:"link"`
		Torrent  string `json:"torrent"`
		Magnet   string `json:"magnet"`
	} `json:"data"`
}
