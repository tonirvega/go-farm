package main

import (
	"github.com/rivo/tview"
)

var (
	cantidadGallinas              = 2
	cantidadTrabajadores          = 1
	huevosPorSegundo              = 1
	totalHuevoJornada             = 0
	jornadaFinalizada             = true
	totalPaquetesJornada          = 0
	jornada                       = 0
	app                           = tview.NewApplication()
	jornadaFinalizadaTrabajadores = make(chan bool)
	jornadaFinalizadaGallinas     = make(chan bool)
	canalHuevos                   = make(chan Huevo, 1000)
	canalPaquetes                 = make(chan Paquete, 1000)
	table                         = tview.NewTable()
)

func main() {
	iniciarFrontend()
}

func iniciarFrontend() {

	mainView := tview.NewFlex().
		AddItem(form(), 0, 1, true).
		AddItem(mainViewComponent(), 0, 2, false)

	go updateDebug()

	if err := app.SetRoot(mainView, true).SetFocus(mainView).Run(); err != nil {

		panic(err)

	}
}

func comenzarJornada() {

	jornada++

	go producirHuevos(canalHuevos)

	go empaquetarHuevos(canalHuevos, canalPaquetes)
}
