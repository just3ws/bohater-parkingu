package models

import (
	"encoding/json"
	"time"

	"github.com/labstack/gommon/log"
)

// RateScheduleEntry model
type RateScheduleEntry struct {
	Days  string `json:"days"`
	Times string `json:"times"`
	Price int32  `json:"price"`
}

// RateSchedule model
type RateSchedule struct {
	Rates []RateScheduleEntry `json:"rates"`
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

	// if err != nil {
	// 	log.Error(err)
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	return 0, nil
}
