package handlers

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jharvs97/2k-blacktop/players"
)

func HandleIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{}, "layouts/main")
}

func HandleGenerate(c *fiber.Ctx) error {

	playerCount, err := strconv.Atoi(c.FormValue("player-count", "2"))

	if err != nil {
		panic(err)
	}

	config, err := players.GetConfigByName(c.FormValue("config", ""), playerCount)

	if err != nil {
		panic(err)
	}

	team1, team2 := generate(config, playerCount)

	fmt.Println(team1)
	fmt.Println(team2)

	return c.Render("partials/players", fiber.Map{"team1": team1, "team2": team2})
}

func HandleUpdateConfig(c *fiber.Ctx) error {
	playerCount, err := strconv.Atoi(c.FormValue("player-count", "2"))

	if err != nil {
		panic(err)
	}

	var configs = players.Configs[playerCount]

	return c.Render("partials/config", fiber.Map{"configs": configs})
}

func HandleDefaultConfig(c *fiber.Ctx) error {
	var configs = players.Configs[2]
	return c.Render("partials/config", fiber.Map{"configs": configs})
}

var r *rand.Rand

func generate(config players.TeamConfig, playerCount int) ([]string, []string) {

	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	teamSize := playerCount / 2

	team1 := make([]string, 0, teamSize)
	team2 := make([]string, 0, teamSize)

	appendPlayersToTeams(&team1, &team2, players.Guards, config.NumGuards)
	appendPlayersToTeams(&team1, &team2, players.Wings, config.NumWings)
	appendPlayersToTeams(&team1, &team2, players.Bigs, config.NumBigs)

	return team1, team2
}

func appendPlayersToTeams(team1 *[]string, team2 *[]string, players []players.Player, numToAppend int) {
	for i := 0; i < numToAppend; i++ {
		i1 := r.Intn(len(players))
		i2 := r.Intn(len(players))

		*team1 = append(*team1, players[i1].Name)
		*team2 = append(*team2, players[i2].Name)
	}
}
