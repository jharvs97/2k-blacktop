package model

type Player struct {
	Name string
	Team string
}

type TeamConfig struct {
	Name      string
	NumGuards int
	NumWings  int
	NumBigs   int
}
