package main

import (
	"fmt"
	"time"
)

func employee(id int, eggs <-chan Egg) {
	for {
		select {

		// Se comprueba si la jornada ha finalizado
		case <-employeeWorkingDayEndChannel:
			debug(fmt.Sprintf(messages.EmployeeEndGoroutine, id))
			return
		default:
			for i := 0; i < 6; i++ {
				select {
				// Se vuelve a comprobar si la jornada ha finalizado incluso cuando el trabajador está empaquetando huevos,
				// digamos que el trabajador está en medio de empaquetar un huevo y la jornada finaliza, el trabajador es un poco
				// radical y deja de empaquetar huevos y se va a su casa.
				case <-employeeWorkingDayEndChannel:
					debug(fmt.Sprintf(messages.EmployeeEndGoroutine, id))
					return
				case <-eggs: // Consumir un huevo del canal
				}
			}

			time.Sleep(1 * time.Second) // Esperar un segundo

			debug(fmt.Sprintf(messages.EmployeePackedEggs, id))

			packagesCountPerWorkingDay++

			if appMode == Terminal {
				updateRow(table)
			}
		}
	}
}

func packEggs(eggs chan Egg) {

	for i := 0; i < employeesAmount; i++ {

		go employee(i, eggs)

	}
}
