package api

import (
	"fmt"

	"backend-config.Cache/config"

	"github.com/gofiber/fiber/v2"
)
//To get the key
func GetKey(c *fiber.Ctx) error {
	key := c.Query("key")

	fmt.Println(" map = ", config.Lru.CacheMap)

	val, ok := config.Lru.Get(key)
	if ok {

		c.Status(200).JSON(&fiber.Map{
			"status":  true,
			"key ":    key,
			"val":     val,
			"message": "Key Fetched successfully",
		})

	} else {
		c.Status(400).JSON(
			&fiber.Map{
				"status":  false,
				"message": "key not found",
			},
		)
	}
	return nil
}
