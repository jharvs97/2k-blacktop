package players

import "fmt"

type TeamConfig struct {
	Name      string
	NumGuards int
	NumWings  int
	NumBigs   int
}

var Configs = map[int][]TeamConfig{
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

func GetConfigByName(name string, playerCount int) (TeamConfig, error) {
	var configs = Configs[playerCount]

	for _, config := range configs {
		if config.Name == name {
			fmt.Println("Found config ", config, " for name ", name)
			return config, nil
		}
	}

	return TeamConfig{}, fmt.Errorf("can't find config '%s' for player count '%d'", name, playerCount)
}
