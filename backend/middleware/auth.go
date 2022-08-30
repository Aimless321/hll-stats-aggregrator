package middleware

import (
	"backend/models"
	"github.com/gofiber/fiber/v2"
	"log"
)

func RequireValidSession(ctx *fiber.Ctx) error {
	sessData, err := models.Store.Get(ctx)
	if err != nil {
		log.Panic(err)
	}

	username := sessData.Get("username")
	if username == nil {
		return ctx.SendStatus(401)
	}

	return ctx.Next()
}
