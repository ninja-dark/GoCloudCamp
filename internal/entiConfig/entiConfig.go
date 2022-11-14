package enticonfig

type Config struct{
	ID int `json:"id"`
	Service string `json:"service"`
	MyData map[string]string `json:"data"`
}






