package bohaterparkingu

import (
	"os"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRate(t *testing.T) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	SetDefaultFailureMode(FailureContinues)
	defer SetDefaultFailureMode(FailureHalts)

	Convey("Rate serialization is case-insensitive", t, func() {
		var starts time.Time
		var ends time.Time

		starts, _ = time.Parse(time.RFC3339, "2016-08-29T08:00:00.000Z")
		ends, _ = time.Parse(time.RFC3339, "2016-08-29T09:00:00.000Z")

		var rate Rate

		Convey("When keys are Capitalized", func() {
			example := `{"Price": 1200, "Starts": "2016-08-29T08:00:00.000Z", "Ends": "2016-08-29T09:00:00.000Z"}`

			rate = UnmarshallRate(&example)

			So(rate.Price, ShouldEqual, 1200)
			So(rate.Starts, ShouldEqual, starts)
			So(rate.Ends, ShouldEqual, ends)
		})

		Convey("When keys are lower-cased", func() {
			example := `{"price": 1200, "starts": "2016-08-29T08:00:00.000Z", "ends": "2016-08-29T09:00:00.000Z"}`

			rate = UnmarshallRate(&example)

			So(rate.Price, ShouldEqual, 1200)
			So(rate.Starts, ShouldEqual, starts)
			So(rate.Ends, ShouldEqual, ends)
		})
	})
}
