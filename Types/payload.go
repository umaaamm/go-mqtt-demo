package Types

type Sensor struct {
	SensorSuhuAir string `json:"sensorSuhuAir" bson:"sensorSuhuAir"`
	SensorSuhu    string `json:"sensorSuhu" bson:"sensorSuhu"`
	SensorPPM     string `json:"sensorPPM" bson:"sensorPPM"`
	SensorPh      string `json:"sensorPh" bson:"sensorPh"`
	LastUpdate    string `json:"lastUpdate" bson:"lastUpdate"`
}
