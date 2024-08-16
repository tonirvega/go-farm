package main

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var debugText string = ""

var debugTextView *tview.TextView = nil

func form() *tview.Form {
	form := tview.NewForm().
		AddTextView(
			"DESCRIPCIÓN",
			"🐔 Este programa simula un sistema de producción y consumo en una granja. Los productores son gallinas que ponen huevos uno por uno. Una vez que un lote de 6 huevos está listo, los consumidores (trabajadores) los recogen y empaquetan. La interfaz muestra en tiempo real la cantidad de huevos producidos, empaquetados y el estado de las gallinas y los trabajadores. ¡Observa cómo la granja opera con precisión y eficiencia!", 50, 14, true, true).
		AddInputField("🐔 nº GALLINAS", strconv.Itoa(cantidadGallinas), 20, nil, func(text string) {
			handleIntEvent(&cantidadGallinas, text)
		}).
		AddInputField("🥚 HUEVOS / SEGUNDO", strconv.Itoa(huevosPorSegundo), 20, nil, func(text string) {
			handleIntEvent(&huevosPorSegundo, text)
		}).
		AddInputField("👨‍🌾 nº TRABAJADORES", strconv.Itoa(cantidadTrabajadores), 20, nil, func(text string) {
			handleIntEvent(&cantidadTrabajadores, text)
		}).
		AddButton("Comenzar jornada", func() {

			if !jornadaFinalizada {
				debug("La jornada ya ha comenzado.")
				return
			}

			comenzarJornada()
		}).
		AddButton("Finalizar jornada", func() {

			if jornadaFinalizada {
				debug("La jornada ya ha finalizado o no ha comenzado.")
				return
			}

			terminarJornada()

		}).
		AddButton("Salir", func() {

			app.Stop()

		})

	form.SetBorder(true).SetTitle("GO FARM").SetTitleAlign(tview.AlignLeft)

	return form
}

func buildMainViewComponent() *tview.Flex {

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(configureTable(), 0, 2, false)

	if debugModeActive() {

		debugTextView = tview.
			NewTextView().
			SetText(debugText)

		flex.
			AddItem(debugTextView, 0, 1, false)

		debugTextView.
			SetBorder(true).
			SetTitle("< DEBUG />").
			SetBorderColor(tcell.ColorYellow)

	}

	return flex
}

func configureTable() *tview.Table {
	table.
		SetBorders(true).
		SetSeparator(tview.Borders.Vertical)

	addHeaderCells(table)

	table.SetBorderPadding(1, 2, 15, 1).
		SetBorder(true)

	return table
}

func addHeaderCells(table *tview.Table) {

	headers := []string{
		"⏲️ JORNADA",
		"🐔 Nº GALLINAS",
		"% HUEVOS POR SEGUNDO",
		"🥚 HUEVOS PRODUCIDOS 🥚",
		"🧑🏻‍🤝‍🧑🏼 Nº TRABAJADORES",
		"📦 PAQUETES EMPAQUETADOS",
	}

	for i := 0; i < len(headers); i++ {
		table.SetCell(0, i, &tview.TableCell{Text: headers[i]})
	}

}

func updateRow(table *tview.Table) {

	table.SetCell(jornada, 0, &tview.TableCell{Text: fmt.Sprintf("%d", jornada)})
	table.SetCell(jornada, 1, &tview.TableCell{Text: fmt.Sprintf("%d", cantidadGallinas)})
	table.SetCell(jornada, 2, &tview.TableCell{Text: fmt.Sprintf("%d", huevosPorSegundo)})
	table.SetCell(jornada, 2, &tview.TableCell{Text: fmt.Sprintf("%d", totalHuevoJornada)})
	table.SetCell(jornada, 3, &tview.TableCell{Text: fmt.Sprintf("%d", cantidadTrabajadores)})
	table.SetCell(jornada, 4, &tview.TableCell{Text: fmt.Sprintf("%d", totalPaquetesJornada)})
}
