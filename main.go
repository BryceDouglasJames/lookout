package main

import (
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	//specify static path
	path, _ := filepath.Abs("./templates/html")

	//serve all layout and schema files
	engine := html.New(path, ".html")

	//satrt fiber instance
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//serve static files
	app.Static("/js", "./templates/js")
	app.Static("/css", "./templates/css")

	//generate routes
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "app")
	})

	log.Fatal(app.Listen(":3000"))
}
