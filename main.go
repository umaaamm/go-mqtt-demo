package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber/v2"

	data "mqtt-demo/main/Types"
	constant "mqtt-demo/main/constant"
	"mqtt-demo/main/database"
	"mqtt-demo/main/handlers"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	var sensor data.Sensor
	err := json.Unmarshal(msg.Payload(), &sensor)
	if err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return
	}
	err = handlers.SaveSensorData(sensor)
	if err != nil {
		fmt.Println("Failed to save to MongoDB:", err)
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

func main() {

	// Init MongoDB
	if err := database.InitMongoDB("mongodb://localhost:27017", "iotDB", "go "); err != nil {
		panic(err)
	}

	go func() {
		app := fiber.New()
		app.Get("/sensor", handlers.GetSensorData)

		if err := app.Listen(":3000"); err != nil {
			panic(err)
		}
	}()

	var broker = constant.MQTT_URL
	var port = 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("client-mqtt-demo")
	opts.SetUsername(constant.MQTT_USERNAME)
	opts.SetPassword(constant.MQTT_PASSWORD)

	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	subscribe(client)
	publish(client)

	client.Disconnect(250)
}

func subscribe(client mqtt.Client) {
	topic := "demo/mqtt/data"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("Failed to subscribe to topic")
		panic(token.Error())
	}
	fmt.Printf("Subscribed to topic: %s", topic)
}

func getRandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func publish(client mqtt.Client) {
	rand.Seed(time.Now().UnixNano())

	for {
		suhuAir := fmt.Sprintf("%.1f", getRandomFloat(26.0, 30.0)) // e.g. 27.3
		suhu := fmt.Sprintf("%.1f", getRandomFloat(27.0, 32.0))    // e.g. 29.4
		ppm := fmt.Sprintf("%.0f", getRandomFloat(300, 700))       // e.g. 420
		ph := fmt.Sprintf("%.1f", getRandomFloat(6.5, 7.5))        // e.g. 7.1

		payload := data.Sensor{
			SensorSuhuAir: suhuAir,
			SensorSuhu:    suhu,
			SensorPPM:     ppm,
			SensorPh:      ph,
			LastUpdate:    time.Now().Format(time.RFC3339),
		}

		jsonBytes, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			continue
		}

		token := client.Publish("demo/mqtt/data", 0, false, jsonBytes)
		token.Wait()
		if token.Error() != nil {
			fmt.Printf("Failed to publish\n")
			panic(token.Error())
		} else {
			fmt.Printf("Published: %s\n", jsonBytes)
		}

		time.Sleep(1 * time.Second)
	}
}
