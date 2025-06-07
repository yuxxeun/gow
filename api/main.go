package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yuxxeun/gow/api/controllers/urlcontroller"
	"github.com/yuxxeun/gow/api/models"
	"github.com/yuxxeun/gow/api/routes"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	routes.SetupRoutes(app)

	api := app.Group("/api")
	urls := api.Group("/urls")

	urls.Get("/", urlcontroller.Index)
	urls.Get("/:id", urlcontroller.Show)
	urls.Post("/", urlcontroller.Create)
	urls.Put("/:id", urlcontroller.Update)
	urls.Delete("/:id", urlcontroller.Delete)

	app.Listen(":8000")
}
