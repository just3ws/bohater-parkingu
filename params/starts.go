package params

import (
	"errors"
	"net/url"
	"time"
)

// StartsParam from URL
func StartsParam(url *url.URL) (time.Time, error) {
	params, ok := url.Query()["starts"]

	if !ok || len(params) < 1 {
		msg := "starts: param is missing"
		return time.Time{}, errors.New(msg)
	}

	starts, err := time.Parse(time.RFC3339, params[0])
	if err != nil {
		msg := "starts: param is invalid"
		return time.Time{}, errors.New(msg)
	}

	return starts, nil
}
