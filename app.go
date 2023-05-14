package main

import (
	"flag"
	"log"

	"teams/database"
	"teams/handlers"
	"teams/middleware"
	"teams/session"

	"github.com/gofiber/fiber/v2"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	session.Start()
	defer session.Close()

	database.Open()
	defer database.Close()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	middleware.AddStandard(app)

	handlers.Add(app)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}
