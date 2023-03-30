package model

type Post struct {
	ID       int    `json:"id"`
	Message  string `json:"message"`
	FileName string `json:"filename"`
	Edited   string `json:"edited"`
	Created  string `json:"created"`
}

type Config struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	PageLength int    `yaml:"pagelength"`
	UserName   string `yaml:"username"`
	Password   string `yaml:"password"`
	CookieKey  string `yaml:"cookiekey"`
	DataPath   string `yaml:"datapath"`
	ConfigPath string
}
