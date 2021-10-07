package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sutantodadang/adopt-me/v1/handler"
)

func main()  {
	app := fiber.New()

	app.Get("/",handler.HelloWorld)

	app.Listen(":5000")

	fmt.Println("Adopt ME")
}