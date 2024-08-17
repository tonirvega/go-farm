package main

import (
	"github.com/rivo/tview"
)

func main() {

	setAppConfigs()

	switch appMode {

	case Wasm:
		startWasmMode()

	case Terminal:
		startTerminalMode()
	}

}

func startTerminalMode() {

	debugImpl = desktopDebug

	mainView := tview.NewFlex().
		AddItem(form(), 80, 1, true).
		AddItem(buildMainViewComponent(), 0, 2, false)

	go refreshView()

	if err := app.SetRoot(mainView, true).SetFocus(mainView).Run(); err != nil {

		panic(err)

	}
}

func startWasmMode() {

	debugImpl = terminalDebug

	startWorkingDay()

	<-employeeWorkingDayEndChannel

}
