package model

type Task struct {
	Title     string `json:"title"`
	CheckMark bool   `json:"checkMark"`
	Id        int    `json:"id"`
}

var Tasks = []Task{
	{Id: 1, Title: "Покупка хлеба", CheckMark: false},
	{Id: 2, Title: "Покупка яиц", CheckMark: false},
	{Id: 3, Title: "Покупка альфредычей", CheckMark: false},
	{Id: 4, Title: "Покупка водки", CheckMark: false},
}
