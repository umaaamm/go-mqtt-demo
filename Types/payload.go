package Types

type Sensor struct {
	SenorSuhuAir string `json:"senorSuhuAir"`
	SenorSuhu    string `json:"senorSuhu"`
	SensorPPM    string `json:"sensorPPM"`
	SensorPh     string `json:"sensorPh"`
	LastUpdate   string `json:"lastUpdate"`
}
