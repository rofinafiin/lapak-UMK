package main

import (
	gcjson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/rofinafiin/lapak-UMK/controller/helper/exception"
)

func main() {
	_ = godotenv.Load(".env")

	// Start a new fiber app
	app := fiber.New(fiber.Config{
		JSONEncoder:  gcjson.Marshal,
		JSONDecoder:  gcjson.Unmarshal,
		ErrorHandler: exception.ErrHandler,
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(compress.New(compress.Config{
		Level: 5,
	}))
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PATCH,DELETE",
	}))
	app.Use(etag.New())

	app.Listen(":3000")
}
