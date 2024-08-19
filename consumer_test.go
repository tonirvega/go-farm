package main

import (
	"fmt"
	"testing"
	"time"
)

func TestEmployeeConsumer(t *testing.T) {
	// Reinitialize the global variables
	newEggsChannel := make(chan Egg)
	employeeWorkingDayEndChannel := make(chan bool)
	employeesAmount = 1
	eggsCountPerWorkingDay = 0
	packagesCountPerWorkingDay = 0

	debugImpl = terminalDebug
	go packEggs(newEggsChannel)

	for i := 0; i < 7; i++ {

		fmt.Printf("Egg %d\n", i)
		newEggsChannel <- Egg{}
		eggsCountPerWorkingDay++
	}

	time.Sleep(10 * time.Second)
	fmt.Printf("Employee working day end\n")
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
