package main

import (
	"fmt"
	"github.com/ermos/dotenv"
	"github.com/ermos/httpcache/internal/pkg/cache"
	"github.com/ermos/httpcache/internal/pkg/controllers"
	"github.com/ermos/httpcache/internal/pkg/logger"
	"github.com/gofiber/fiber/v2"
	cache2 "github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/utils"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	_ = dotenv.Parse(".env")

	err := dotenv.Require(
		"APP_PORT",
		"LOG_PATH",
		"PUBLIC_PATH",
		"MIN_MINUTES",
		"MAX_MINUTES",
	)
	if err != nil {
		log.Fatal(err)
	}

	logger.Init()
	cache.Init()

	app := fiber.New()

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	app.Use(cache2.New(cache2.Config{
		ExpirationGenerator: func(c *fiber.Ctx, cfg *cache2.Config) time.Duration {
			newCacheTime, _ := strconv.Atoi(c.GetRespHeader("Cache-Time", "600"))
			return time.Second * time.Duration(newCacheTime)
		},
		KeyGenerator: func(c *fiber.Ctx) string {
			return utils.CopyString(c.Path())
		},
	}))

	app.Static("/", os.Getenv("PUBLIC_PATH"))
	app.Get("/:minutes/*", controllers.GetWrapper)

	_ = app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
