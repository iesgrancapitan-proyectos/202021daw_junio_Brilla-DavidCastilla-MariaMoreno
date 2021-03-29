package models

type User struct {
	Username   string `json:"_key,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	Name       string `json:"name,omitempty"`
	Bio        string `json:"bio,omitempty"`
	Birthday   int64  `json:"birthday,omitempty"`
	ProfileImg string `json:"profile_img,omitempty"`
	CreatedAt  int64  `json:"created_at,omitempty"`
	UpdateAt   int64  `json:"update_at,omitempty"`
}
