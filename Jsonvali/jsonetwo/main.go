package main

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name        string `json:"name"`
	Capacity    int    `json:"capacity"`
	Time        string `json:"time"`
	Information Info   `json:"info"`
}

type Info struct {
	Description string `json:"des"`
}

func main() {
	fmt.Println("Hello world")
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2020-03-22T20:07:28Z", "info": {
		"des": "a sensor reading"
	}}`
	var reading SensorReading
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)
}
