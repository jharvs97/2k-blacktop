package handlers

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jharvs97/2k-blacktop/database"
	"github.com/jharvs97/2k-blacktop/model"
)

func HandleIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{}, "layouts/main")
}

func HandleGenerateTeam(c *fiber.Ctx) error {
	configName := c.FormValue("config", "Guards")
	config, err := database.GetConfigByName(configName)
	if err != nil {
		panic(err)
	}

	team1, team2 := generateTeams(config)

	return c.Render("partials/players", fiber.Map{"team1": team1, "team2": team2})
}

func HandleUpdateConfig(c *fiber.Ctx) error {
	playerCount, err := strconv.Atoi(c.FormValue("player-count", "2"))
	if err != nil {
		panic(err)
	}

	configs, _ := database.GetConfigsByPlayerCount(playerCount)

	return c.Render("partials/config", fiber.Map{"configs": configs})
}

func HandleDefaultConfig(c *fiber.Ctx) error {
	configs, _ := database.GetConfigsByPlayerCount(2)
	return c.Render("partials/config", fiber.Map{"configs": configs})
}

var r *rand.Rand

func generateTeams(config model.TeamConfig) ([]string, []string) {
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	teamSize := config.PlayerCount / 2

	team1 := make([]string, 0, teamSize)
	team2 := make([]string, 0, teamSize)

	guards, _ := database.GetNRandomPlayers(config.NumGuards, model.Guard)
	wings, _ := database.GetNRandomPlayers(config.NumWings, model.Wing)
	bigs, _ := database.GetNRandomPlayers(config.NumBigs, model.Big)

	appendPlayersToTeams(&team1, &team2, guards)
	appendPlayersToTeams(&team1, &team2, wings)
	appendPlayersToTeams(&team1, &team2, bigs)

	return team1, team2
}

func appendPlayersToTeams(team1 *[]string, team2 *[]string, players []model.Player) {
	for _, p := range players[0 : len(players)/2] {
		*team1 = append(*team1, p.Name)
	}

	for _, p := range players[len(players)/2:] {
		*team2 = append(*team2, p.Name)
	}
}
