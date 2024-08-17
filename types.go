package main

type Egg struct{}

type RunMode string

const (
	Wasm     RunMode = "WASM"
	Terminal RunMode = "TERMINAL"
)

var toRunMode = map[string]RunMode{
	"TERMINAL": Terminal,
	"WASM":     Wasm,
}

type Messages struct {
	AppModeIsNotSet           string `json:"APP_MODE_IS_NOT_SET"`
	AppModeIsInvalid          string `json:"APP_MODE_IS_INVALID"`
	HenLaidAnEgg              string `json:"HEN_LAID_AN_EGG"`
	HenEndGoroutine           string `json:"HEN_END_GOROUTINE"`
	EmployeeEndGoroutine      string `json:"EMPLOYEE_END_GOROUTINE"`
	EmployeePackedEggs        string `json:"EMPLOYEE_PACKED_EGGS"`
	WorkingDayIsOver          string `json:"WORKING_DAY_IS_OVER"`
	Description               string `json:"DESCRIPTION"`
	DescriptionTitle          string `json:"DESCRIPTION_TITLE"`
	NumChickens               string `json:"NUM_CHICKENS"`
	EggsPerSecond             string `json:"EGGS_PER_SECOND"`
	NumWorkers                string `json:"NUM_WORKERS"`
	StartWorkingDay           string `json:"START_WORKING_DAY"`
	WorkingDayAlreadyStarted  string `json:"WORKING_DAY_ALREADY_STARTED"`
	FinishWorkingDay          string `json:"FINISH_WORKING_DAY"`
	WorkingDayAlreadyFinished string `json:"WORKING_DAY_ALREADY_FINISHED"`
	Exit                      string `json:"EXIT"`
	GoFarm                    string `json:"GO_FARM"`
	Debug                     string `json:"DEBUG"`
	WorkingDay                string `json:"WORKING_DAY"`
	NumChickensHeader         string `json:"NUM_CHICKENS_HEADER"`
	EggsPerSecondHeader       string `json:"EGGS_PER_SECOND_HEADER"`
	EggsProduced              string `json:"EGGS_PRODUCED"`
	NumWorkersHeader          string `json:"NUM_WORKERS_HEADER"`
	PackagesPacked            string `json:"PACKAGES_PACKED"`
}
