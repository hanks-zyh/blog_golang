package model

type User struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Pwd       string `json:"-"`
	Phone     string `json:"phone"`
	AuthData  Auth   `json:authData`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

type Auth struct {
	Token    string `json:"-"`
	Platform string `json:"platform"`
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
}
