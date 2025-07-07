package main

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"

	"mqtt-demo/main/Types"
	"mqtt-demo/main/database"
)

func TestMessagePubHandler_SaveToMongoDB(t *testing.T) {
	// Init DB (gunakan DB khusus test)
	err := database.InitMongoDB("mongodb://localhost:27017", "test_iotDB", "sensor")
	assert.NoError(t, err)

	// Dummy sensor
	payload := Types.Sensor{
		SensorSuhuAir: "28.0",
		SensorSuhu:    "29.5",
		SensorPPM:     "500",
		SensorPh:      "7.1",
		LastUpdate:    time.Now().Format(time.RFC3339),
	}

	jsonData, _ := json.Marshal(payload)

	// Simulasikan message dari MQTT
	msg := mockMessage{topic: "demo/mqtt/data", payload: jsonData}
	messagePubHandler(nil, msg)

	// Tunggu insert selesai
	time.Sleep(100 * time.Millisecond)

	// Cek apakah data tersimpan
	var result Types.Sensor
	err = database.SensorCollection.FindOne(context.TODO(), bson.M{"sensorSuhu": "29.5"}).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, "29.5", result.SensorSuhu)
}

type mockMessage struct {
	topic   string
	payload []byte
}

func (m mockMessage) Duplicate() bool   { return false }
func (m mockMessage) Qos() byte         { return 0 }
func (m mockMessage) Retained() bool    { return false }
func (m mockMessage) Topic() string     { return m.topic }
func (m mockMessage) MessageID() uint16 { return 1 }
func (m mockMessage) Payload() []byte   { return m.payload }
func (m mockMessage) Ack()              {}
