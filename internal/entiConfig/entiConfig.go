package enticonfig

type Config struct{
	ID int `json:"id"`
	Service string `json:"service"`
	Data []MyData `json:"message"`
}

type MyData struct{
	ID int `json:"id"`
	Key string `json:"key"`
	Value string `json:"value"`
	IDConf int `json:"idconf"`
}





