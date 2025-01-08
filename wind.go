package aviation

import (
	"fmt"
	"regexp"
)

// 22016KT or 12012MPS
const windSpeedPattern = "(?P<direction>[0-9]{3})(?P<speed>[0-9]{2,3})(?P<units>KT|MPS)"

// 22010G22KT
const windSpeedGustingPattern = "(?P<direction>[0-9]{3})(?P<speed>[0-9]{2,3})G(?P<gust>[0-9]{2,3})(?P<units>KT|MPS)"

// 22010KT 190V260
const windSpeedVariablePattern = "(?P<direction>[0-9]{3})(?P<speed>[0-9]{2,3})(?P<units>KT|MPS) (?P<from>[0-9]{3})V(?P<to>[0-9]{3})"

// 22010G22KT 190V260
const windSpeedGustingVariablePattern = "(?P<direction>[0-9]{3})(?P<speed>[0-9]{2,3})G(?P<gust>[0-9]{2,3})(?P<units>KT|MPS) (?P<from>[0-9]{3})V(?P<to>[0-9]{3})"

type Wind struct {
	direction    string
	speed        string
	gust         string
	units        string
	variableFrom string
	variableTo   string
}

func parseWind(message string) (*Wind, error) {

	windSpeedRegexp := regexp.MustCompile(windSpeedPattern)
	windSpeedGustingRegexp := regexp.MustCompile(windSpeedGustingPattern)
	windSpeedVariableRegexp := regexp.MustCompile(windSpeedVariablePattern)
	windSpeedGustingVariableRegexp := regexp.MustCompile(windSpeedGustingVariablePattern)

	wind := &Wind{}
	if windSpeedGustingVariableRegexp.MatchString(message) {
		result := getGroups(windSpeedGustingVariableRegexp, message)
		wind.direction = result["direction"]
		wind.speed = result["speed"]
		wind.units = result["units"]
		wind.gust = result["gust"]
		wind.variableFrom = result["from"]
		wind.variableTo = result["to"]
	} else if windSpeedVariableRegexp.MatchString(message) {
		result := getGroups(windSpeedVariableRegexp, message)
		wind.direction = result["direction"]
		wind.speed = result["speed"]
		wind.units = result["units"]
		wind.gust = result["gust"]
		wind.variableFrom = result["from"]
		wind.variableTo = result["to"]
	} else if windSpeedGustingRegexp.MatchString(message) {
		result := getGroups(windSpeedGustingRegexp, message)
		wind.direction = result["direction"]
		wind.speed = result["speed"]
		wind.units = result["units"]
		wind.gust = result["gust"]
	} else if windSpeedRegexp.MatchString(message) {
		result := getGroups(windSpeedRegexp, message)
		wind.direction = result["direction"]
		wind.speed = result["speed"]
		wind.units = result["units"]
	} else {
		return nil, fmt.Errorf("no wind found in metar message '%v'", message)
	}

	return wind, nil
}

func getGroups(re *regexp.Regexp, message string) map[string]string {
	result := make(map[string]string)
	m := re.FindStringSubmatch(message)

	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = m[i]
		}
	}
	return result
}
