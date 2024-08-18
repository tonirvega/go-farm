package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"slices"
	"strconv"
)

// App config
var (
	appMode                = Terminal
	debugMode              = true
	debugImpl func(string) = nil
	messages  Messages     = Messages{}
	lang      string       = "en"
)

func setAppConfigs() {

	mode, err := getAppMode()

	if err != nil {

		panic(err)

	}

	appMode = toRunMode[mode]

	switch appMode {

	case Wasm:

		debugImpl = terminalDebug

		loadLangFromMemory()

	case Terminal:

		debugImpl = desktopDebug

		loadLangFromJsonFile()

	}

	debugMode = debugModeActive()

}

func loadLangFromJsonFile() {
	langStruct, err := readLangFile()

	if err != nil {
		panic(err)

	}

	messages = *langStruct
}

func loadLangFromMemory() {

	messages = Messages{
		AppModeIsNotSet:           "APP_MODE is not set, please set it to 'WASM' or 'TERMINAL'",
		AppModeIsInvalid:          "APP_MODE %s is invalid, please set it to 'WASM' or 'TERMINAL'",
		HenLaidAnEgg:              "The hen %d laid an egg",
		HenEndGoroutine:           "End goroutine for hen %d",
		EmployeeEndGoroutine:      "End goroutine for employee %d",
		EmployeePackedEggs:        "Employee %d packed 6 eggs",
		WorkingDayIsOver:          "Working day is over. You can start a new day now",
		Description:               "🐔 This program simulates a production and consumption system on a farm. The producers are hens that lay eggs one by one. Once a batch of 6 eggs is ready, the consumers (workers) collect and pack them. The interface shows in real-time the number of eggs produced, packed, and the status of the hens and workers. Watch how the farm operates with precision and efficiency!",
		DescriptionTitle:          "DESCRIPTION",
		NumChickens:               "🐔 Hens",
		EggsPerSecond:             "🥚 Eggs / sec",
		NumWorkers:                "👨‍🌾 Employees",
		StartWorkingDay:           "Start Working Day",
		WorkingDayAlreadyStarted:  "The working day has already started.",
		FinishWorkingDay:          "Finish Working Day",
		WorkingDayAlreadyFinished: "The working day has already finished or has not started.",
		Exit:                      "Exit",
		GoFarm:                    "Go Farm",
		Debug:                     "< Debug mode />",
		WorkingDay:                "⏲️ Working Day",
		NumChickensHeader:         "🐔 Hens amount",
		EggsPerSecondHeader:       "🥚 Eggs / sec",
		EggsProduced:              "🥚 Produced eggs",
		NumWorkersHeader:          "🧑🏻‍🤝‍🧑🏼 Workers",
		PackagesPacked:            "📦 Packages Packed",
	}
}

func debugModeActive() bool {

	debugModeActive := false

	if value := os.Getenv("DEBUG_MODE"); value != "" {
		if parsedBool, err := strconv.ParseBool(value); err == nil {
			debugModeActive = parsedBool
		} else {
			panic(err)
		}
	}

	return debugModeActive
}

func getAppMode() (string, error) {

	envMode := os.Getenv("APP_MODE")

	if envMode == "" {

		fmt.Printf(messages.AppModeIsNotSet)

		return "WASM", nil
	}

	if slices.Contains([]string{string(Wasm), string(Terminal)}, envMode) {
		return envMode, nil
	}

	return "", fmt.Errorf("Error, %s is not a valid app mode", envMode)
}

func readLangFile() (*Messages, error) {

	file, err := os.Open(fmt.Sprintf("lang.%s.json", lang))

	if err != nil {
		return nil, err
	}

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	var messages Messages
	if err := json.Unmarshal(byteValue, &messages); err != nil {
		return nil, err
	}

	return &messages, nil
}
