package main

import "github.com/rivo/tview"

// Glboal variables
var (
	hensAmount                   = 2
	employeesAmount              = 1
	eggsPerSecond                = 1
	eggsCountPerWorkingDay       = 0
	packagesCountPerWorkingDay   = 0
	workingDay                   = 0
	workingDayIsOver             = true
	employeeWorkingDayEndChannel = make(chan bool)
	chickenWorkingDayEndChannel  = make(chan bool)
	eggsChannel                  = make(chan Egg, 1000)
	app                          = tview.NewApplication()
	table                        = tview.NewTable()
)
