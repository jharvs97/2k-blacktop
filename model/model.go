package model

type Player struct {
	Name     string
	Team     string
	Position Position
}

type TeamConfig struct {
	Name        string
	NumGuards   int
	NumWings    int
	NumBigs     int
	PlayerCount int
}

type Position int

const (
	Guard Position = 1
	Wing  Position = 2
	Big   Position = 3
)

type Teams struct {
	Team1 []Player
	Team2 []Player
}
