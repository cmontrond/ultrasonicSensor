package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	g "gobot.io/x/gobot/platforms/dexter/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)


func robotRunLoop(gopigo3 *g.Driver, ultrasonicSensor *aio.AnalogSensorDriver) {
	for {

		ultrasonicSensorVal, ultrasonicSensorErr := ultrasonicSensor.Read()

		if ultrasonicSensorErr != nil {
			fmt.Errorf("Error reading sensor %+v", ultrasonicSensorErr)
		}

		fmt.Println("Sensor Value is ", ultrasonicSensorVal)


		time.Sleep(time.Second)
	}
}

func main() {
	raspiAdaptor := raspi.NewAdaptor()
	gopigo3 := g.NewDriver(raspiAdaptor)

	ultrasonicSensor := aio.NewAnalogSensorDriver(gopigo3, "AD_2_1")

	mainRobotFunc := func() {
		robotRunLoop(gopigo3, ultrasonicSensor)
	}

	robot := gobot.NewRobot("gopigo3sensorChecker",
		[]gobot.Connection{raspiAdaptor},
		[]gobot.Device{gopigo3, ultrasonicSensor},
		mainRobotFunc,
	)

	robot.Start()
}
