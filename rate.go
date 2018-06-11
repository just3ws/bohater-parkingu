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

// UnmarshallRate from JSON string
func UnmarshallRate(str *string) Rate {
	var rate Rate
	err := json.Unmarshal([]byte(*str), &rate)
	if err != nil {
		panic(err)
	}

	return rate
}
