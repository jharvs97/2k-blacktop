package handlers

import (
	"strconv"

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

	teams := generateTeams(config)

	return c.Render("partials/teams", fiber.Map{"teams": teams})
}

func HandleGenerateConfig(c *fiber.Ctx) error {
	playerCount, err := strconv.Atoi(c.FormValue("player-count", "2"))
	if err != nil {
		panic(err)
	}

	configs, _ := database.GetConfigsByPlayerCount(playerCount)

	return c.Render("partials/config", fiber.Map{"configs": configs})
}

func generateTeams(config model.TeamConfig) model.Teams {
	teams, err := database.GetTeamsForConfig(config)
	if err != nil {
		panic(err)
	}

	return teams
}
