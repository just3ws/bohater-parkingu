package models

import (
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
)

// RateSchedule model
type RateSchedule struct {
	Rates []RateScheduleEntry `json:"rates"`
}

// RateScheduleEntry model
type RateScheduleEntry struct {
	Days  string `json:"days"`
	Times string `json:"times"`
	Price int32  `json:"price"`
}

type RateRange struct {
	Hours []int32
	Price int32
}

func (rr *RateRange) Include(hour int) bool {
	log.Debug(hour)
	return false
}

// UnmarshalRateSchedule from JSON
func UnmarshalRateSchedule(str *string) RateSchedule {
	var rateSchedule RateSchedule
	err := json.Unmarshal([]byte(*str), &rateSchedule)
	if err != nil {
		panic(err)
	}

	return rateSchedule
}

// RateScheduleLookup for range of dates
func RateScheduleLookup(starts time.Time, ends time.Time) (int32, error) {

	log.Debug(starts)
	log.Debug(ends)

	return 0, nil
}
