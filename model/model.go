package model

type Post struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	Edited  string `json:"edited"`
	Created string `json:"created"`
}

type Config struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	ApiDir     string `yaml:"apidir"`
	PageLength int    `yaml:"pagelength"`
	UserName   string `yaml:"username"`
	Password   string `yaml:"password"`
	CookieKey  string `yamml:"cookiekey"`
}
