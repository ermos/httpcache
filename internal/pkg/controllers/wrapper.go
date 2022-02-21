package controllers

import (
	"fmt"
	"github.com/ermos/httpcache/internal/pkg/response"
	"github.com/ermos/httpcache/internal/pkg/wrapper"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	minMinutes = 0
	maxMinutes = 0
)

func GetWrapper(c *fiber.Ctx) error {
	url := strings.TrimLeft(c.OriginalURL(), "/"+c.Params("minutes")+"/")

	minutes, err := strconv.Atoi(c.Params("minutes"))
	if err != nil {
		return c.Status(400).JSON(response.Error{
			Code:    1,
			Message: "invalid format for minute's parameter",
		})
	}

	if minMinutes == 0 {
		minMinutes, err = strconv.Atoi(os.Getenv("MIN_MINUTES"))
		if err != nil {
			log.Fatal(err)
		}
	}

	if maxMinutes == 0 {
		maxMinutes, err = strconv.Atoi(os.Getenv("MAX_MINUTES"))
		if err != nil {
			log.Fatal(err)
		}
	}

	if minutes < minMinutes {
		return c.Status(400).JSON(response.Error{
			Code:    1,
			Message: fmt.Sprintf("minimum cache time is %s minutes", os.Getenv("MIN_MINUTES")),
		})
	}

	if minutes > maxMinutes {
		return c.Status(400).JSON(response.Error{
			Code:    2,
			Message: fmt.Sprintf("maximum cache time is %s minutes", os.Getenv("MAX_MINUTES")),
		})
	}

	expAt := time.Now().Add(time.Minute * time.Duration(minutes)).Unix()

	err = wrapper.Handle(c, url, expAt)
	if err != nil {
		log.Println(err)
		return c.Status(400).JSON(response.Error{
			Code:    3,
			Message: "we cannot access to resource and no cache found for it",
		})
	}

	return nil
}
