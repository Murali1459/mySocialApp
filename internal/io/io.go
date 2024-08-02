package io

type User struct {
	Id         int    `json:"id"`
	UserName   string `json:"userName"`
	Email      string `json:"email"`
	ProfilePic string `json:"profie_pic_url"`
	Bio        string `json:"bio"`
}

type RegisterUser struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Post struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id"`
	Content  string `json:"content"`
	ImageUrl string `json:"image_url"`
}

type Followers struct {
	User User `json:"user"`
}
