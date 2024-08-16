package main

import "strconv"

func comenzarJornada() {

	jornadaFinalizada = false
	jornada++

	jornadaFinalizadaTrabajadores = make(chan bool)
	jornadaFinalizadaGallinas = make(chan bool)
	canalHuevos = make(chan Huevo, 1000)

	go producirHuevos(canalHuevos)
	go empaquetarHuevos(canalHuevos)
}

func terminarJornada() {

	for i := 0; i < cantidadGallinas; i++ {
		jornadaFinalizadaGallinas <- true
	}

	for i := 0; i < cantidadTrabajadores; i++ {
		jornadaFinalizadaTrabajadores <- true
	}

	totalPaquetesJornada = 0
	totalHuevoJornada = 0
	jornadaFinalizada = true

	debug("Jornada en curso finalizada. Puedes iniciar una nueva jornada.")
}

func handleIntEvent(number *int, text string) {

	parsedInt, err := strconv.Atoi(text)

	if err != nil {
		*number = 0
	}

	*number = parsedInt

}
