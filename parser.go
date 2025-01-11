package aviation

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func parse(message string, metar *Metar) error {

	log.Printf("parse metar message: '%v'\n", message)

	messageParts := strings.Split(message, " ")

	pos := 0

	// icao location
	metar.Location = messageParts[pos]
	pos++

	// time of the report
	t, err := time.Parse("021504Z200601", messageParts[pos]+time.Now().Format("200601"))
	if err != nil {
		return fmt.Errorf("error parsing time of the observation %v", err)
	}
	metar.DateTime = t
	metar.DateTimeString = messageParts[pos]
	pos++

	// is it an automated report
	if messageParts[pos] == "AUTO" {
		metar.IsAuto = true
		pos++
	} else {
		metar.IsAuto = false
	}

	// the runway wind
	metar.Wind = parseWind(messageParts[pos])
	pos++

	// wind variability
	// !! the variabilty VRB20KT or VRB10MPS will not match but is also possible in metar message !!
	if len(messageParts[pos]) == 7 && messageParts[pos][3] == 'V' {
		intFrom, _ := strconv.Atoi(messageParts[pos][:3])
		intTo, _ := strconv.Atoi(messageParts[pos][4:])
		metar.Wind.VariableFrom = uint(intFrom)
		metar.Wind.VariableTo = uint(intTo)
		pos++
	}

	// visibility
	metar.Visibility = strings.TrimLeft(messageParts[pos], "0")
	pos++

	// runway visibility
	for repeat := true; repeat; repeat = (messageParts[pos][0] == 'R' && messageParts[pos][3] == '/') {
		vis := parseRwyVisibility(messageParts[pos])
		if vis.Runway != "" {
			metar.RwyVisibility = append(metar.RwyVisibility, vis)
		}
		pos++
	}

	// weather
	metar.Weather = parseWeather(messageParts[pos])
	pos++
	return nil
}

func parseWeather(in string) Weather {
	log.Printf("parse weather: %v\n", in)
	weatherRegexp := regexp.MustCompile(`(?P<intensity>\\+|-)|(?P<vercinity>VC)?(?<weather>(BZ|BL|BR|DR|DS|DU|DZ|FG|FC|FU|FZ|GR|GS|HZ|IC|MI|PL|PO|RA|SA|SG|SH|SN|SQ|SS|TS|VA|UP|RE){1,3})`)
	weather := Weather{}
	if !weatherRegexp.MatchString(in) {
		log.Printf("weather string '%v' does not match weather regexp\n", in)
		return weather
	}

	result := getGroups(weatherRegexp, in)
	weather.Intensity = result["intensity"]
	if result["vercinity"] == "VC" {
		weather.Vicinity = true
	}
	if result["weather"] == "" {
		return weather
	}

	for i := 0; i < len(result["weather"]); i = i + 2 {
		weather.Phenomena = append(weather.Phenomena, result["weather"][i:i+2])
	}

	return weather
}

func parseRwyVisibility(in string) RunwayVisibility {
	log.Printf("parse runway visibility: %v\n", in)
	rwyVisibilityRegexp := regexp.MustCompile(`(?<runway>\d{2}[RLC]?)/(?<evolution>[MP]?)(?<visibility>\d{4})(?<tendency>[DUN]?)(V(?<variablevisibility>\d{4})(?<variabletendency>[DUN]))?`)

	rwyVis := RunwayVisibility{}
	if !rwyVisibilityRegexp.MatchString(in) {
		return rwyVis
	}

	result := getGroups(rwyVisibilityRegexp, in)
	rwyVis.Runway = result["runway"]
	intVis, _ := strconv.Atoi(result["visibility"])
	rwyVis.Visibility = uint(intVis)
	rwyVis.Evolution = result["evolution"]
	rwyVis.Tendency = result["tendency"]
	intVariable, _ := strconv.Atoi(result["variablevisibility"])
	rwyVis.Variable = uint(intVariable)
	rwyVis.VariableTendency = result["variabletendency"]
	return rwyVis
}

func parseWind(in string) Wind {
	log.Printf("parse wind: %v\n", in)
	windSpeedRegexp := regexp.MustCompile(`(?P<direction>\d{3})(?P<speed>\d{2,3})(G(?P<gust>\d{2}))?(?P<units>KT|MPS)`)
	w := Wind{}
	if !windSpeedRegexp.MatchString(in) {
		return w
	}

	result := getGroups(windSpeedRegexp, in)
	intDirection, _ := strconv.Atoi(result["direction"])
	w.Direction = uint(intDirection)

	intSpeed, _ := strconv.Atoi(result["speed"])
	w.Speed = uint(intSpeed)

	w.Units = result["units"]

	if result["gust"] != "" {
		intGust, _ := strconv.Atoi(result["gust"])
		w.Gust = uint(intGust)
	}
	return w
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
