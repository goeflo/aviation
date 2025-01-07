package aviation

import (
	"fmt"
	"regexp"
)

const icaoPattern = "([A-Z]{4})"

var icaoRegex = regexp.MustCompile(icaoPattern)

type Location struct {
	icao string
}

func parseLocation(message string) (*Location, error) {
	if !icaoRegex.MatchString(message) {
		return nil, fmt.Errorf("no ICAO station id found in metar message")
	}
	icao := icaoRegex.FindAllStringSubmatch(message, -1)[0][1]
	return &Location{icao: icao}, nil
}
