package main

import (
	"fmt"
	"testing"
	"time"
)

func TestEmployeeConsumer(t *testing.T) {
	// Reinitialize the global variables
	eggsChannel := make(chan Egg)
	employeeWorkingDayEndChannel := make(chan bool)

	eggsCountPerWorkingDay = 0
	packagesCountPerWorkingDay = 0

	debugImpl = terminalDebug
	go packEggs(eggsChannel)

	for i := 0; i < 7; i++ {

		eggsChannel <- Egg{}
		eggsCountPerWorkingDay++
	}

	time.Sleep(10 * time.Second)
	employeeWorkingDayEndChannel <- true

	if packagesCountPerWorkingDay != 1 {

		errorMessage := fmt.Sprintf("Employee does not package correctly, expected 1 package count, got %d", packagesCountPerWorkingDay)
		t.Fatal(errorMessage)
	}

	if eggsCountPerWorkingDay != 7 {

		errorMessage := fmt.Sprintf("Employee does not package correctly, expected 7 eggs count, got %d", eggsCountPerWorkingDay)

		t.Fatal(errorMessage)

	}
}
