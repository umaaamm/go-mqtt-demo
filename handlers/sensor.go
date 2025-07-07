package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"

	"mqtt-demo/main/Types"
	"mqtt-demo/main/database"
)

func SaveSensorData(sensor Types.Sensor) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := database.SensorCollection.InsertOne(ctx, sensor)
	return err
}

func GetSensorData(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.SensorCollection.Find(ctx, fiber.Map{})
	if err != nil {
		return c.Status(500).SendString("Failed to query MongoDB")
	}
	defer cursor.Close(ctx)

	var results []Types.Sensor
	for cursor.Next(ctx) {
		var doc Types.Sensor
		if err := cursor.Decode(&doc); err != nil {
			return c.Status(500).SendString("Failed to decode MongoDB document")
		}
		results = append(results, doc)
	}

	return c.JSON(results)
}
