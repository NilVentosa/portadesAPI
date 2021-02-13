package portades

type Portada struct {
	Id        int    `json:"id"`
	Intro     string `json:"intro"`
	Newspaper string `json:"newspaper"`
	Headline  string `json:"headline"`
	Result    bool   `json:"result"`
	Video     string `json:"video"`
	Episode   string `json:"episode"`
}
