package dto

type User struct {
	ChatId   int64  `json:"chat_id"`
	UserName string `json:"username" gorm:"column:username"`
}

type UsersLinks struct {
	ChatId   int64  `json:"chat_id"`
	Link     string `json:"link" gorm:"column:irl"`
	Category string `json:"category"`
	Filters  string `json:"filters"`
}
