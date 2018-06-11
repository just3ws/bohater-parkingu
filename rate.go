package bohaterparkingu

import (
	"encoding/json"
	"time"
)

// Rate model
type Rate struct {
	Price  int32     `json:"price"`
	Starts time.Time `json:"starts"`
	Ends   time.Time `json:"ends"`
}

// UnmarshalRate from JSON string
func UnmarshalRate(str *string) Rate {
	var rate Rate
	err := json.Unmarshal([]byte(*str), &rate)
	if err != nil {
		panic(err)
	}

	return rate
}

// MarshalRate to JSON string
func MarshalRate(rate *Rate) string {
	b, err := json.Marshal(rate)
	if err != nil {
		panic(err)
	}
	return string(b)
}
