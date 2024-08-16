package handlers

import (
    "fiber-app-paafff/internal/domain/services"
    "github.com/gofiber/fiber/v2"
    "strconv"
)

type UserHandler struct {
    userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
    }

    user, err := h.userService.GetUser(uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).SendString("User not found")
    }

    return c.JSON(user)
}