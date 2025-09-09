package router

import (
	"github.com/dennyaris/portal-news/internal/delivery/http/handler"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, uh *handler.UserHandler, ch *handler.CategoryHandler, nh *handler.ContentHandler) {
	app.Get("/healthz", func(c *fiber.Ctx) error { return c.SendString("ok") })

	v1 := app.Group("/api/v1")

	users := v1.Group("/users")
	users.Post("/", uh.Create)
	users.Get("/:id", uh.Get)
	users.Get("/", uh.List)

	cats := v1.Group("/categories")
	cats.Post("/", ch.Create)
	cats.Get("/:id", ch.Get)
	cats.Get("/", ch.List)
	cats.Put("/:id", ch.Update)
	cats.Delete("/:id", ch.Delete)

	news := v1.Group("/contents")
	news.Post("/", nh.Create)
	news.Get("/:id", nh.Get)
	news.Get("/", nh.List)
	news.Put("/:id", nh.Update)
	news.Delete("/:id", nh.Delete)
}
