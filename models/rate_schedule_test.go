package models

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"

	"time"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRateSchedule(t *testing.T) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	SetDefaultFailureMode(FailureContinues)
	defer SetDefaultFailureMode(FailureHalts)

	var starts time.Time

	starts, _ = time.Parse(time.RFC3339, "2016-08-29T08:00:00.000Z")

	Convey("UnmarshalRateSchedule", t, func() {
		log.Debug(starts.Weekday())

		var rateSchedule RateSchedule

		Convey("Read the rate schedule data", func() {
			file, err := ioutil.ReadFile("../data/rate_schedule.json")
			if err != nil {
				panic(err)
			}

			example := string(file)

			rateSchedule = UnmarshalRateSchedule(&example)

			for i := 0; i < len(rateSchedule.Rates); i++ {
				log.Debug("{{{")
				log.Debug(rateSchedule.Rates[i])
				shortDays := strings.Split(rateSchedule.Rates[i].Days, ",")
				log.Debug(shortDays)

				days := make([]string, len(shortDays))
				for i, day := range shortDays {
					days[i] = strings.ToUpper(day)
				}

				log.Debug(days)

				times := strings.Split(rateSchedule.Rates[i].Times, "-")
				log.Debug(times)

				hours := make([]int32, len(times))
				for i, v := range times {
					x64, _ := strconv.Atoi(v)
					hours[i] = int32(x64)
				}

				log.Debug(hours)

				log.Debug(rateSchedule.Rates[i].Price)

				log.Debug("}}}")

			}
		})
	})
}
