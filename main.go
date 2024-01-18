package main

import (
	"log"
	"me/pickside/common"
	"me/pickside/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	err := run()

	if err != nil {
		panic(err)
	}
}

func Connect(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func run() error {
	err := common.InitDB()

	if err != nil {
		return err
	}

	defer common.CloseDB()

	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		AppName:      "Pickside me-service v1.0.1",
	})

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Post("/connect", router.Connect)
	v1.Post("/disconnect", router.Disconnect)
	v1.Post("/refresh-token", router.RefreshToken)
	v1.Post("/register", router.Register)

	log.Fatal(app.Listen(":8080"))

	return nil
}
