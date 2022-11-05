package enticonfig

type Config struct{
	Id int `json:"id"`
	Service string `json:"service"`
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
	Version int `json:"version"`
}

