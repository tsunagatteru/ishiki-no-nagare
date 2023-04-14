package model

type Post struct {
	ID       int    `json:"id"`
	Message  string `json:"message"`
	FileName string `json:"filename"`
	Edited   string `json:"edited"`
	Created  string `json:"created"`
}
