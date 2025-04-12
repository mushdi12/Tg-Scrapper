package dto

import "fmt"

type Tuser struct {
	Username string `json:"username"`
}

func Print(ms string) {
	fmt.Println(ms)
}
