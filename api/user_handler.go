package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roblesdotdev/hotel-reservation/db"
	"github.com/roblesdotdev/hotel-reservation/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	u, err := h.userStore.GetUserById(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(u)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users := h.userStore.GetAllUsers(c.Context())
	return c.JSON(users)
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.UserInputParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	u, err := h.userStore.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(u)
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.userStore.DeleteUser(c.Context(), id); err != nil {
		return err
	}
	return c.JSON("User deleted")
}

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	return c.JSON("TODO")
}
