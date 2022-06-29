package model

type Character struct {
	Name  string
	level []*Level
}

type Level struct {
	Number uint
}
