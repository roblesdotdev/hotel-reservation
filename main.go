package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/roblesdotdev/hotel-reservation/api"
)

func main() {
	port := flag.String("port", "8080", "The listen port of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUserById)
	app.Listen(fmt.Sprintf(":%s", *port))
}
