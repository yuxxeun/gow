package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/yuxxeun/gow/api/controllers/urlcontroller"
	"github.com/yuxxeun/gow/api/models"
	"github.com/yuxxeun/gow/api/routes"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	app.Use(cors.New())

	routes.SetupRoutes(app)

	api := app.Group("/api")
	urls := api.Group("/urls")

	urls.Get("/", urlcontroller.Index)
	urls.Post("/", urlcontroller.Create)
	urls.Get("/:id", urlcontroller.Show)
	urls.Patch("/:id", urlcontroller.Update)
	urls.Delete("/:id", urlcontroller.Delete)

	app.Listen(":8000")
}
