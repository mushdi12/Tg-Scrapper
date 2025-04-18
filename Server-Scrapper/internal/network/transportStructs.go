package network

type Client struct {
	ChatId   int64  `json:"chat_id" gorm:"column:chat_id"`
	UserName string `json:"username" gorm:"column:username"`
}

type ClientLink struct {
	ChatId   int64  `json:"chat_id"`
	Link     string `json:"link" gorm:"column:irl"`
	Category string `json:"category"`
	Filters  string `json:"filters"`
}

type GetLink struct {
	ChatId int64 `json:"chat_id"`
}

type Links struct {
	Link []string `json:"links"`
}

type RemoveLink struct {
	ChatId int64  `json:"chat_id"`
	Link   string `json:"link" gorm:"column:irl"`
}
