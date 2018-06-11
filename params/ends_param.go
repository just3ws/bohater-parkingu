package params

import (
	"errors"
	"net/url"
	"time"
)

// EndsParam from URL
func EndsParam(url *url.URL) (time.Time, error) {
	params, ok := url.Query()["ends"]

	if !ok || len(params) < 1 {
		msg := "ends: param is missing"
		return time.Time{}, errors.New(msg)
	}

	ends, err := time.Parse(time.RFC3339, params[0])
	if err != nil {
		msg := "ends: param is invalid"
		return time.Time{}, errors.New(msg)
	}

	return ends, nil
}
