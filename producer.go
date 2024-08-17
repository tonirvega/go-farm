package main

import (
	"fmt"
	"time"
)

func hen(id int, eggsPerSecond int, eggs chan<- Egg) {

	for {
		select {
		case <-chickenWorkingDayEndChannel:
			debug(fmt.Sprintf(messages.HenEndGoroutine, id))
			time.Sleep(50 * time.Millisecond) // Esperar un segundo
			return
		default:
			// producir huevos
			for i := 0; i < eggsPerSecond; i++ {
				eggs <- Egg{}
				eggsCountPerWorkingDay++
				debug(fmt.Sprintf(messages.HenLaidAnEgg, id))
			}

			time.Sleep(1 * time.Second) // Esperar un segundo
		}
	}
}

func produceEggs(eggs chan Egg) {

	for i := 0; i < chickensAmount; i++ {

		go hen(i, eggsPerSecond, eggs) // Cada gallina produce n huevos por segundo

	}

}
