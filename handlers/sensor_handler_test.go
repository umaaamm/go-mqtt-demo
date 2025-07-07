package handlers

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"mqtt-demo/main/Types"
	"mqtt-demo/main/database"
)

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic("Failed to connect to MongoDB: " + err.Error())
	}
	database.SensorCollection = client.Database("testdb").Collection("sensors")
}

func TestSaveSensorData(t *testing.T) {
	t.Parallel()

	sensor := Types.Sensor{
		SensorSuhuAir: "27.5",
		SensorSuhu:    "28.1",
		SensorPPM:     "400",
		SensorPh:      "6.8",
		LastUpdate:    time.Now().Format(time.RFC3339),
	}

	err := SaveSensorData(sensor)
	assert.NoError(t, err)

	var result Types.Sensor
	err = database.SensorCollection.FindOne(context.TODO(), bson.M{"sensorSuhu": "28.1"}).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, "28.1", result.SensorSuhu)
}

func TestGetSensorData(t *testing.T) {
	t.Parallel()

	app := fiber.New()
	app.Get("/sensors", GetSensorData)

	// Insert dummy data
	_ = SaveSensorData(Types.Sensor{
		SensorSuhuAir: "26.5",
		SensorSuhu:    "27.0",
		SensorPPM:     "300",
		SensorPh:      "7.0",
		LastUpdate:    time.Now().Format(time.RFC3339),
	})

	// Send request
	req, err := http.NewRequest("GET", "/sensors", nil)
	assert.NoError(t, err)

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
