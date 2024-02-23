package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/roblesdotdev/hotel-reservation/api"
	"github.com/roblesdotdev/hotel-reservation/db"
	"github.com/roblesdotdev/hotel-reservation/types"
)

func main() {
	port := flag.String("port", "8080", "The listen port of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	memoryUserStore := db.NewMemoryUserStore([]*types.User{
		{
			ID:        "001",
			FirstName: "John",
			LastName:  "Doe",
		},
		{
			ID:        "002",
			FirstName: "Mary",
			LastName:  "Jane",
		},
	})
	userHandler := api.NewUserHandler(memoryUserStore)

	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUserById)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Put("/user/:id", userHandler.HandlePutUser)
	app.Listen(fmt.Sprintf(":%s", *port))
}
