package main

type User struct {
	Id int `json:"id"`
	UserName string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Users []User
