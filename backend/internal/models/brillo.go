package models

type Brillo struct {
	Content   string   `json:"content"`
	Media     []string `json:"media"`
	CreatedAt int64    `json:"created_at"`
}
