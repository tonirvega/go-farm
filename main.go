package main

import (
	"github.com/rivo/tview"
)

var debugImpl func(string) = terminalDebug

func main() {

	if modoDesktopActivo() {

		// wasm no es compatible con tview
		mostrarPanel()

	} else {

		comenzarJornada()

		<-jornadaFinalizadaTrabajadores

	}

}

func mostrarPanel() {

	debugImpl = desktopDebug

	mainView := tview.NewFlex().
		AddItem(form(), 0, 1, true).
		AddItem(buildMainViewComponent(), 0, 2, false)

	go updateDebug()

	if err := app.SetRoot(mainView, true).SetFocus(mainView).Run(); err != nil {

		panic(err)

	}
}
