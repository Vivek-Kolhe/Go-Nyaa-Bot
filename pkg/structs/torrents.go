package structs

type Torrents struct {
	Count int `json:"count"`
	Data  []struct {
		Title    string `json:"title"`
		Link     string `json:"link"`
		Size     string `json:"size"`
		Seeders  int    `json:"seeders"`
		Leechers int    `json:"leechers"`
	} `json:"data"`
}
