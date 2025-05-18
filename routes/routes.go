package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuxxeun/gow/server/controllers/urlcontroller"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	urls := v1.Group("/urls")

	urls.Get("/", urlcontroller.Index)
	urls.Get("/:id", urlcontroller.Show)
	urls.Post("/", urlcontroller.Create)
	urls.Patch("/:id", urlcontroller.Update)
	urls.Delete("/:id", urlcontroller.Delete)
}
