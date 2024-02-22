package api

import "github.com/gofiber/fiber/v2"

func HandleGetUsers(ctx *fiber.Ctx) error {
	return ctx.JSON("USERS")
}

func HandleGetUserById(ctx *fiber.Ctx) error {
	return ctx.JSON("USER BY ID")
}
