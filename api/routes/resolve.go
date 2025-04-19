package routes

import (
	"github.com/adix/shortapi/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveUrl(c *fiber.Ctx) error {
	url := c.Params("url")
	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short url not found",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "connection to redis failed",
		})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()
	//Increments the value of the key "counter" in Redis database 1 by 1.
	_ = rInr.Incr(database.Ctx, "counter")
	return c.Redirect(value, 301)

}
