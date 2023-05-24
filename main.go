package main

import (
	"github.com/aiteung/musik"
	gcjson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	_ "github.com/rofinafiin/lapak-UMK/controller/docs"
	"github.com/rofinafiin/lapak-UMK/controller/helper/exception"
	"github.com/rofinafiin/lapak-UMK/url"
)

type ReturnData struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

// @title LAPAK-UMK OPEN API PROTOTYPE
// @version 2.0

// @host localhost:3000
// @BasePath /
// @schemes https
// @Failure      400  {object}	ReturnData
// @Failure      404  {object}  ReturnData
// @Failure      500  {object}  ReturnData
func main() {
	_ = godotenv.Load("prod.env")

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

	// Setup the router
	url.SetupRoutes(app)
	app.Get("/docs/*", swagger.HandlerDefault)

	// Listen on PORT 3000
	app.Listen(musik.Dangdut())
}
