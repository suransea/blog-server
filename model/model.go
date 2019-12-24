package model

type Article struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
