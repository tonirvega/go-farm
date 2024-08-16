package main

import (
	"fmt"
	"time"
)

func trabajador(id int, eggs <-chan Huevo) {
	for {
		select {

		// Se comprueba si la jornada ha finalizado
		case <-jornadaFinalizadaTrabajadores:
			debug(fmt.Sprintf("Fin go routine worker %d", id))
			return
		default:
			for i := 0; i < 6; i++ {
				select {
				// Se vuelve a comprobar si la jornada ha finalizado incluso cuando el trabajador está empaquetando huevos,
				// digamos que el trabajador está en medio de empaquetar un huevo y la jornada finaliza, el trabajador es un poco
				// radical y deja de empaquetar huevos y se va a su casa.
				case <-jornadaFinalizadaTrabajadores:
					debug(fmt.Sprintf("Fin go routine trabajador %d", id))
					return
				case <-eggs: // Consumir un huevo del canal
				}
			}

			time.Sleep(1 * time.Second) // Esperar un segundo

			debug(fmt.Sprintf("Trabajador %d empaquetó 6 huevos.", id))

			totalPaquetesJornada++

			if modoDesktopActivo() {
				updateRow(table)
			}
		}
	}
}

func empaquetarHuevos(eggs chan Huevo) {

	for i := 0; i < cantidadTrabajadores; i++ {

		go trabajador(i, eggs)

	}
}
