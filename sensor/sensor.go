package rpi2cube

import (
  "regexp"
  "os"
	"bufio"
  "strconv"
  "log"
)

type Sensor struct {
  Name string
  Temp float64

}

func AllSensors() (sensors []Sensor) {
  sensorsFile, _ := os.Open("/sys/bus/w1/devices/w1_bus_master1/w1_master_slaves")
	sensor_files := bufio.NewScanner(sensorsFile)

	for sensor_files.Scan() {
		sensors = append(sensors, NewSensor(sensor_files.Text()) )
	}

	return sensors
}

func NewSensor(name string) (sensor Sensor) {
  sensor.Name = name
  return sensor
}

func (sensor Sensor)File() (* os.File) {
  file, err :=  os.Open("/sys/bus/w1/devices/"+sensor.Name+"/w1_slave")
  if(err != nil) {
		log.Print(err)
	}
	return file
}

func (sensor Sensor)ReadSensor() (Sensor){
	tempFinder := regexp.MustCompile(`t=(-?\d+)`)

  sensorFile := sensor.File()
	sensorScanner := bufio.NewScanner(sensorFile)

	// advance to second line
	sensorScanner.Scan()
	sensorScanner.Scan()

	tempReading := tempFinder.FindStringSubmatch(sensorScanner.Text())[1]
	temp, _ := strconv.ParseFloat(tempReading, 32)
	sensor.Temp = temp / 1000

  return sensor
}
