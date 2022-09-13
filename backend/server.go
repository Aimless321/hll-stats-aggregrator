package main

import (
	"backend/handlers"
	"backend/middleware"
	"backend/models"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3"
	"github.com/gookit/config/v2"
	jsonDriver "github.com/gookit/config/v2/json"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/discord"
	"github.com/shareed2k/goth_fiber"
	"log"
	"time"
)

func main() {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(jsonDriver.Driver)

	err := config.LoadFiles("config.json")
	if err != nil {
		panic(err)
	}

	models.DbPool, err = pgxpool.Connect(context.Background(), config.String("databaseUrl"))
	if err != nil {
		log.Panicf("Unable to connect to database: %v\n", err)
	}
	defer models.DbPool.Close()

	storage := sqlite3.New()
	models.Store = session.New(session.Config{
		Storage:      storage,
		Expiration:   30 * 24 * time.Hour,
		CookieSecure: true,
	})
	models.Store.RegisterType("")

	app := fiber.New()
	app.Use(cors.New(
		cors.Config{
			Next:             nil,
			AllowOrigins:     "*",
			AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
			AllowHeaders:     "",
			AllowCredentials: true,
			ExposeHeaders:    "",
			MaxAge:           0,
		},
	))
	app.Use(logger.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	goth.UseProviders(
		discord.New(
			config.String("clientId"),
			config.String("clientSecret"),
			config.String("discordCallback"),
			discord.ScopeIdentify,
		),
	)

	app.Get("/auth/:provider", goth_fiber.BeginAuthHandler)
	app.Get("/auth/:provider/callback", handlers.DiscordCallback)
	app.Get("/logout", handlers.Logout)
	app.Get("/stats/recruitment", middleware.RequireValidSession, handlers.GetRecruitmentData)
	app.Get("/stats/:steamid", middleware.RequireValidSession, handlers.GetData)

	if err := app.Listen(":8080"); err != nil {
		log.Panic(err)
	}
}
