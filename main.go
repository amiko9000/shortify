package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"shortify/config"
	"shortify/database"
	"shortify/router"
)

func Init(app *fiber.App) {
	config.Config()
	database.Connect()
	router.Register(app)
}

func main() {
	app := fiber.New()
	Init(app)
	log.Fatal(app.Listen(":3000"))
}