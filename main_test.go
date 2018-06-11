package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	SetDefaultFailureMode(FailureContinues)
	defer SetDefaultFailureMode(FailureHalts)
}

type thing struct {
	a string
}

func panics() {
	panic("Goofy Gophers!")
}

const timeLayout = "2006-01-02 15:04"
