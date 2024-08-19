package main

import (
	"testing"
)

func TestEventIntHandler(t *testing.T) {

	intRef := 0

	handleIntEvent(&intRef, "2")

	if intRef != 2 {

		t.Fatal("Int ref not equals to right value")

	}

}
