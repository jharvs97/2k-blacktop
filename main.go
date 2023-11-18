package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django/v3"
	"github.com/jharvs97/2k-blacktop/database"
	"github.com/jharvs97/2k-blacktop/handlers"
)

func main() {
	err := database.Init()

	if err != nil {
		panic(err)
	}

	engine := django.New("./views", ".html")

	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	app.Get("/", handlers.HandleIndex)
	app.Get("/generateConfig", handlers.HandleGenerateConfig)
	app.Get("/generateTeam", handlers.HandleGenerateTeam)

	addr := ":6969"

	err = app.Listen(addr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening on ", addr)
}
