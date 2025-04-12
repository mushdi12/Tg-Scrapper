package dto

type Client struct {
	ChatId   int64  `json:"chat_id"`
	UserName string `json:"user_name"`
	Link     string `json:"link"`
	Category string `json:"category"`
	Filter   string `json:"filter"`
}
