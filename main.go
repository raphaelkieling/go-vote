package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/raphaelkieling/go-vote/config"
	"github.com/raphaelkieling/go-vote/handlers"
)

func main() {
	db := config.ConnectToDatabase()
	config := config.NewConfig()

	app := fiber.New()

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			config.Auth.Username: config.Auth.Password,
		},
	}))

	// Using cors
	app.Use(cors.New())

	apiGroup := app.Group("/api")

	campaignHandler := handlers.CampaignHandler{
		DB: db,
	}

	apiGroup.Get("/campaign", campaignHandler.GetAll)
	apiGroup.Post("/campaign/:id/vote", campaignHandler.Vote)
	apiGroup.Post("/campaign", campaignHandler.Create)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Port)))
}
