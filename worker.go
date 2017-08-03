package main

type Params struct {
	Source      string `json:"src"`
	Destination string `json:"dst"`
	Width       int    `json:"width"`
	Heigth      int    `json:"height"`
}
