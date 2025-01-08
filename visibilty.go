package aviation

import (
	"fmt"
	"regexp"
	"strings"
)

/*
	runway visibilty examples:

R26/0400 = RVR runway 26 with range 400 meters
R25/M0075 = RVR runway 25 is less than 75 meters (M=Minus)
R33L/P1500 = RVR runway 33 LEFT is greater than 1500 meters (P=Plus)
R16R/1000D = RVR runway 16 RIGHT is 1000 meters with aggravation (D=Down)
R16R/1000U = RVR runway 16 RIGHT is 1000 meters with improvement (U=UP)
R33C/0900N = RVR runway 33 CENTRE is 900 meters with no change (N=No change)
R27/0150V0300U = RVR runway 27 variable (V) from 150 to 300 meters with improvement (U=Up)
*/

const visibiltyPattern = "( [0-9]{4} | CAVOK | NSC )"
const visibilityAllRunwaysPattern = "R\\d{2}[RCL]?\\/[MP]?\\d{4}[DUN]?(?:V\\d{4}[DUN]?)?"
const visibilitySingleRunwayPattern = "(?<rwy>R\\d{2}[RCL]?)\\/(?<evo>[MP]?)(?<vis>\\d{4})(?<bcm>[DUN]?)(?:V(?<variable>\\d{4})(?<vbcm>[DUN]?))?"

var visibiltyRegex = regexp.MustCompile(visibiltyPattern)
var visibilityAllRunwaysRegex = regexp.MustCompile(visibilityAllRunwaysPattern)
var visibilitySingleRunwaysRegex = regexp.MustCompile(visibilitySingleRunwayPattern)

type runwayVisualRange struct {
	rwy            string
	visibiltyRange string
	becoming       string
	evolution      string
	variableRange  string
	variableBcm    string
}
type Visibility struct {
	visibiltyRange string
	runway         []runwayVisualRange
}

func parseVisibility(message string) (*Visibility, error) {

	vis := &Visibility{}
	visRange, err := parseRange(message)
	if err != nil {
		vis.visibiltyRange = "????"
	}

	vis.visibiltyRange = visRange

	rwyVis := parseRunwayVisibilty(message)
	if rwyVis != nil {
		vis.runway = append(vis.runway, rwyVis...)
	}

	return vis, nil
}

func parseRunwayVisibilty(message string) []runwayVisualRange {

	if visibilityAllRunwaysRegex.MatchString(message) {
		rwyVis := visibilityAllRunwaysRegex.FindAllStringSubmatch(message, -1)
		return parseRunwaysString(rwyVis)
	}

	return nil
}

func parseRunwaysString(runways [][]string) []runwayVisualRange {
	rwyVisRange := []runwayVisualRange{}

	for _, r := range runways {
		result := getGroups(visibilitySingleRunwaysRegex, r[0])
		rwyVisRange = append(rwyVisRange, runwayVisualRange{
			rwy:            result["rwy"],
			evolution:      result["evo"],
			becoming:       result["bcm"],
			visibiltyRange: result["vis"],
			variableRange:  result["variable"],
			variableBcm:    result["vbcm"],
		})
	}
	return rwyVisRange
}

func parseRange(message string) (string, error) {
	if !visibiltyRegex.MatchString(message) {
		return "", fmt.Errorf("no visibility details found in metar message")
	}
	vis := strings.Trim(visibiltyRegex.FindAllStringSubmatch(message, -1)[0][1], " ")
	return vis, nil
}
