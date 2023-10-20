package main

import (
	"log"

	"github.com/RestApi/restApi/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	//Routes and End points
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
