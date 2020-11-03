package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2020-03-22T20:07:28Z", "info": {
		"des": "a sensor reading"
	}}`
	var reading map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)
}
