package main

import (
	"github.com/rivo/tview"
)

func main() {

	mainView := tview.NewFlex().
		AddItem(form(), 0, 1, true).
		AddItem(buildMainViewComponent(), 0, 2, false)

	go updateDebug()

	if err := app.SetRoot(mainView, true).SetFocus(mainView).Run(); err != nil {

		panic(err)

	}

}
