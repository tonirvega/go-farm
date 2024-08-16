package main

import (
	"fmt"
	"time"
)

func gallina(id int, eggsPerSecond int, eggs chan<- Huevo) {

	for {
		select {
		case <-jornadaFinalizadaGallinas:
			debug(fmt.Sprintf("Fin go routine de producción de huevos para la gallina -> %d", id))
			time.Sleep(50 * time.Millisecond) // Esperar un segundo
			return
		default:
			// producir huevos
			for i := 0; i < eggsPerSecond; i++ {
				eggs <- Huevo{}
				totalHuevoJornada++
				debug(fmt.Sprintf("Gallina %d puso un huevo.", id))
			}

			time.Sleep(1 * time.Second) // Esperar un segundo
		}
	}
}

func producirHuevos(eggs chan Huevo) {

	for i := 0; i < cantidadGallinas; i++ {

		go gallina(i, huevosPorSegundo, eggs) // Cada gallina produce n huevos por segundo

	}

}
