package database

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/jharvs97/2k-blacktop/model"
)

var (
	guards []model.Player
	wings  []model.Player
	bigs   []model.Player
)

var configs = map[int][]model.TeamConfig{
	2: {
		{
			Name:      "Guards",
			NumGuards: 1,
		},
		{
			Name:     "Wings",
			NumWings: 1,
		},
		{
			Name:    "Bigs",
			NumBigs: 1,
		},
	},
	4: {
		{
			Name:      "Two Guards",
			NumGuards: 2,
		},
		{
			Name:      "Guard and Wing",
			NumGuards: 1,
			NumWings:  1,
		},
		{
			Name:      "Guard and Big",
			NumGuards: 1,
			NumBigs:   1,
		},
		{
			Name:     "Two Wings",
			NumWings: 2,
		},
		{
			Name:     "Wing and Big",
			NumWings: 1,
			NumBigs:  1,
		},
	},
	6: {
		{
			Name:      "One Guard, One Wing and One Big",
			NumGuards: 1,
			NumWings:  1,
			NumBigs:   1,
		},
		{
			Name:      "3 Guards",
			NumGuards: 3,
		},
	},
	10: {
		{
			Name:      "Standard 5v5",
			NumGuards: 1,
			NumWings:  2,
			NumBigs:   2,
		},
	},
}

func GetConfigByName(name string, playerCount int) (model.TeamConfig, error) {
	var configs = configs[playerCount]

	for _, config := range configs {
		if config.Name == name {
			fmt.Println("Found config ", config, " for name ", name)
			return config, nil
		}
	}

	return model.TeamConfig{}, fmt.Errorf("can't find config '%s' for player count '%d'", name, playerCount)
}

func GetConfigsByPlayerCount(playerCount int) ([]model.TeamConfig, error) {
	return configs[playerCount], nil
}

func GetAllGuards() []model.Player {
	return guards
}

func GetAllWings() []model.Player {
	return guards
}

func GetAllBigs() []model.Player {
	return guards
}

func Init() error {
	file, err := os.Open("./database/2015-2016.csv")

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

		p := model.Player{Name: name, Team: team}

		switch pos {
		case "PG":
			fallthrough
		case "PG-SG":
			guards = append(guards, p)
		case "SG":
			fallthrough
		case "SF":
			fallthrough
		case "SG-SF":
			fallthrough
		case "SF-PF":
			wings = append(wings, p)
		case "PF":
			fallthrough
		case "PF-C":
			fallthrough
		case "C":
			bigs = append(bigs, p)
		default:
			fmt.Printf("Couldn't match %s to a position\n", pos)
		}
	}

	return nil
}
