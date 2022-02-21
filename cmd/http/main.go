package main

import (
	"fmt"
	"github.com/ermos/dotenv"
	"github.com/ermos/httpcache/internal/pkg/cache"
	"github.com/ermos/httpcache/internal/pkg/controllers"
	"github.com/ermos/httpcache/internal/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	_ = dotenv.Parse(".env")

	err := dotenv.Require(
		"APP_PORT",
		"LOG_PATH",
		"MIN_MINUTES",
		"MAX_MINUTES",
	)
	if err != nil {
		log.Fatal(err)
	}

	logger.Init()
	cache.Init()

	app := fiber.New()

	app.Static("/", "./public")
	app.Get("/:minutes/*", controllers.GetWrapper)

	_ = app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
