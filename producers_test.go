package main

import (
	"testing"
	"time"
)

func TestConsumersAndProducers(t *testing.T) {

	// Global variables
	debugImpl = terminalDebug
	chickensAmount = 2

	// run 6 times in order to insert 6 items in the channel, it should, produce
	// one package
	startWorkingDay()

	// wait until the producers finish their work
	time.Sleep(8 * time.Second)

	// wait 8 seconds for the goroutines to finish
	// 2 packages is ok, because there are two producers, 12 egs, which equals to 2 packages.
	if packagesCountPerWorkingDay != 2 {
		t.Errorf("Expected 1 package, got %d", packagesCountPerWorkingDay)
	}

}
