package model

type Post struct{
	ID int `json:"id"`
	Message string `json:"message"`
	Edited string `json:"edited"`
	Created string `json:"created"`
}
