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
		AddDropDown(
			"LANG",
			[]string{"en", "es"}, 0, func(option string, index int) {

				lang = option

				loadLangFromJsonFile()
			},
		).
		AddTextView(
			messages.DescriptionTitle,
			messages.Description,
			50,
			14,
			true,
			true).
		AddInputField(
			messages.NumChickensHeader,
			strconv.Itoa(hensAmount),
			20,
			nil,
			// Event handler
			func(text string) {

				handleIntEvent(&hensAmount, text)

			},
		).
		AddInputField(
			messages.EggsPerSecondHeader,
			strconv.Itoa(eggsPerSecond),
			20,
			nil,
			// Event handler
			func(text string) {

				handleIntEvent(&eggsPerSecond, text)

			},
		).
		AddInputField(
			messages.NumWorkersHeader,
			strconv.Itoa(employeesAmount),
			20,
			nil,
			// Event handler
			func(text string) {

				handleIntEvent(&employeesAmount, text)

			},
		).
		AddButton(
			messages.StartWorkingDay,
			// Event handler
			func() {

				if !workingDayIsOver {
					debug(messages.WorkingDayAlreadyStarted)
					return
				}

				startWorkingDay()
			},
		).
		AddButton(
			messages.FinishWorkingDay,
			// Event handler
			func() {

				if workingDayIsOver {
					debug(messages.WorkingDayAlreadyFinished)
					return
				}

				finishWorkingDay()

			},
		).
		AddButton(
			messages.Exit,
			func() {

				app.Stop()

			},
		)

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
			SetTitle(messages.Debug).
			SetBorderColor(tcell.ColorYellow)

	}

	return flex
}

func configureTable() *tview.Table {
	table.
		SetBorders(true).
		SetSeparator(tview.Borders.Vertical)

	addHeaderCells(table)

	table.SetBorderPadding(1, 2, 5, 1).
		SetBorder(true)

	return table
}

func addHeaderCells(table *tview.Table) {

	headers := []string{
		messages.WorkingDay,
		messages.NumChickensHeader,
		messages.EggsPerSecondHeader,
		messages.EggsProduced,
		messages.NumWorkersHeader,
		messages.PackagesPacked,
	}

	for i := 0; i < len(headers); i++ {
		table.SetCell(0, i, &tview.TableCell{Text: headers[i]})
	}

}

func updateRow(table *tview.Table) {

	table.SetCell(workingDay, 0, &tview.TableCell{Text: calculateWorkingDayDuration()})
	table.SetCell(workingDay, 1, &tview.TableCell{Text: fmt.Sprintf("%d", hensAmount)})
	table.SetCell(workingDay, 2, &tview.TableCell{Text: fmt.Sprintf("%d", eggsPerSecond)})
	table.SetCell(workingDay, 3, &tview.TableCell{Text: fmt.Sprintf("%d", eggsCountPerWorkingDay)})
	table.SetCell(workingDay, 4, &tview.TableCell{Text: fmt.Sprintf("%d", employeesAmount)})
	table.SetCell(workingDay, 5, &tview.TableCell{Text: fmt.Sprintf("%d", packagesCountPerWorkingDay)})
}

func calculateWorkingDayDuration() string {

	return timeNow.Sub(timeStart).String()

}
