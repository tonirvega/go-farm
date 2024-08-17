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
	messages  *Messages    = nil
	lang      string       = "en"
)

func setAppConfigs() {

	loadLang()

	mode, err := getAppMode()

	if err != nil {

		panic(err)

	}

	appMode = toRunMode[mode]

	debugMode = debugModeActive()

}

func loadLang() {

	langStruct, err := readLangFile()

	if err != nil {
		panic(err)

	}

	messages = langStruct
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

	return "", fmt.Errorf(messages.AppModeIsInvalid, envMode)
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
