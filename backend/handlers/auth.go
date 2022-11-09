package handlers

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/config/v2"
	"github.com/shareed2k/goth_fiber"
	"log"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func DiscordCallback(ctx *fiber.Ctx) error {
	user, err := goth_fiber.CompleteUserAuth(ctx)
	if err != nil {
		log.Panic(err)
	}

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.Set("accept", "application/json")
	req.Header.Set("authorization", fmt.Sprintf("Bearer %v", user.AccessToken))
	req.SetRequestURI(fmt.Sprintf("https://discord.com/api/users/@me/guilds/%v/member", config.String("guildId")))
	if err := a.Parse(); err != nil {
		log.Panic(err)
	}

	t := models.DiscordGuildMember{}
	var body []byte
	var errs []error
	if _, body, errs = a.Bytes(); len(errs) > 0 {
		log.Panic(errs)
	}
	json.Unmarshal(body, &t)

	if sess, err := models.Store.Get(ctx); err == nil {
		defer sess.Save()
		if !contains(t.Roles, config.String("requiredRole")) {
			return ctx.SendStatus(401)
		}
		sess.Set("username", t.User.Username)
		sess.Set("token", user.AccessToken)

		return ctx.Redirect(fmt.Sprintf("%v?username=%v", config.String("frontendUrl"), t.User.Username))
	}

	return ctx.SendStatus(401)
}

func Logout(ctx *fiber.Ctx) error {
	if err := goth_fiber.Logout(ctx); err != nil {
		log.Panic(err)
	}

	sess, err := models.Store.Get(ctx)
	if err != nil {
		log.Panic(err)
	}
	defer sess.Save()

	sess.Destroy()

	return ctx.SendStatus(201)
}
