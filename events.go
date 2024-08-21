package main

import (
	"strconv"
	"time"
)

var timeStart time.Time = time.Now()
var timeNow time.Time = time.Now()

func startWorkingDay() {

	timeStart = time.Now()
	workingDayIsOver = false
	workingDay++

	employeeWorkingDayEndChannel = make(chan bool)
	chickenWorkingDayEndChannel = make(chan bool)
	eggsChannel = make(chan Egg, 1000)

	go produceEggs(eggsChannel)
	go packEggs(eggsChannel)
}

func finishWorkingDay() {

	for i := 0; i < hensAmount; i++ {
		chickenWorkingDayEndChannel <- true
	}

	for i := 0; i < employeesAmount; i++ {
		employeeWorkingDayEndChannel <- true
	}

	close(eggsChannel)

	packagesCountPerWorkingDay = 0
	eggsCountPerWorkingDay = 0
	workingDayIsOver = true
	timeNow = time.Now()
	debug(messages.WorkingDayIsOver)
}

func handleIntEvent(number *int, text string) {

	parsedInt, err := strconv.Atoi(text)

	if err != nil {
		*number = 0
	}

	*number = parsedInt

}
