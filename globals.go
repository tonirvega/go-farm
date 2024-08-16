package main

import "github.com/rivo/tview"

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
	table                         = tview.NewTable()
)
