package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	SetDefaultFailureMode(FailureContinues)
	defer SetDefaultFailureMode(FailureHalts)

	Convey("Equality assertions should be accessible", t, func() {
		thing1a := thing{a: "asdf"}
		thing1b := thing{a: "asdf"}
		thing2 := thing{a: "qwer"}

		So(1, ShouldEqual, 1)
		So(1, ShouldNotEqual, 2)
		So(1, ShouldAlmostEqual, 1.000000000000001)
		So(1, ShouldNotAlmostEqual, 2, 0.5)
		So(thing1a, ShouldResemble, thing1b)
		So(thing1a, ShouldNotResemble, thing2)
		So(&thing1a, ShouldPointTo, &thing1a)
		So(&thing1a, ShouldNotPointTo, &thing1b)
		So(nil, ShouldBeNil)
		So(1, ShouldNotBeNil)
		So(true, ShouldBeTrue)
		So(false, ShouldBeFalse)
		So(0, ShouldBeZeroValue)
	})
}

type thing struct {
	a string
}

func panics() {
	panic("Goofy Gophers!")
}

const timeLayout = "2006-01-02 15:04"
