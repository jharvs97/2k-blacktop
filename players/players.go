package players

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Player struct {
	Name string
	Team string
}

var (
	Guards []Player
	Wings  []Player
	Bigs   []Player
)

func Init() error {
	file, err := os.Open("./data/2015-2016.csv")

	if err != nil {
		return err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()

	if err != nil {
		return err
	}

	for _, row := range data[1:] {

		name := row[1]
		pos := row[2]
		team := row[4]

		p := Player{Name: name, Team: team}

		switch pos {
		case "PG":
			fallthrough
		case "PG-SG":
			Guards = append(Guards, p)
		case "SG":
			fallthrough
		case "SF":
			fallthrough
		case "SG-SF":
			fallthrough
		case "SF-PF":
			Wings = append(Wings, p)
		case "PF":
			fallthrough
		case "PF-C":
			fallthrough
		case "C":
			Bigs = append(Bigs, p)
		default:
			fmt.Printf("Couldn't match %s to a position\n", pos)
		}
	}

	return nil
}
