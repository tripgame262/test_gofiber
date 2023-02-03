package main

import (
	"fmt"
	"gobasic/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDb()
	app := fiber.New()
	db := database.Database

	app.Get("/api/sum", func(c *fiber.Ctx) error {
		sumType := c.Query("sum")

		var sumValue float64

		db.Db.Table("power").Select("SUM(" + sumType + ")").Row().Scan(&sumValue)

		fmt.Print(sumValue)
		return c.JSON(fiber.Map{
			"type": sumType,
			"sum":  sumValue,
		})
	})

	app.Listen(":3050")

}
