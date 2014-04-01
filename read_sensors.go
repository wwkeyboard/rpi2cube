package main

import (
	"fmt"
	"encoding/json"
	"os"
	"bytes"
  "net/http"
	"time"
	"github.com/wwkeyboard/rpi2cube/sensor"
)

type SensorMessage struct {
  Sensor string `json:"sensor"`
  C      float64
}

type CubeMessage struct {
  Type string        `json:"type"`
  Time string        `json:"time"`
  Data SensorMessage `json:"data"`
}

func main() {
	fmt.Printf("These devices:\n")

	for _,sensor := range rpi2cube.AllSensors() {
		s := sensor.ReadSensor()

	  now := time.Now().Local().Format(time.RFC3339Nano)

		d := CubeMessage{"temp", now,
              SensorMessage{s.Name, s.Temp}}

		arr_d := make([]CubeMessage, 1,1)
		arr_d[0] = d

		payload, _ := json.Marshal(arr_d)

		_, err := http.Post("http://192.168.1.13:8180/1.0/event/put",
			"application/json",
      bytes.NewReader(payload))
		if err != nil {fmt.Printf(err.Error())}

		os.Stdout.Write(payload)
	}
}
