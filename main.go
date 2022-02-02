package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/api/v1/news", proxy.Forward("http://localhost:10001/api/v1/news"))

	// app.Listen(":8080")

	app.ListenTLS(":8080", "./cert/server.crt", "./cert/server.key")

}
