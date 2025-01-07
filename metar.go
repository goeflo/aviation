package aviation

import (
	"fmt"
	"strings"
)

/*
https://www.metar-taf-decoder.com/
https://aviationweather.gov/data/metar/?decoded=1&ids=EDDB
https://en.wikipedia.org/wiki/METAR
https://www.seabirds.de/ato/Weather-METAR.php
https://wiki.ivao.aero/en/home/training/documentation/METAR_explanation
https://met.nps.edu/~bcreasey/mr3222/files/helpful/DecodeMETAR-TAF.html
*/
type Metar struct {
	Message  string
	Location string
	DateTime string
	IsAuto   bool
	Wind     *Wind
}

// new metar struct
func NewMetar() *Metar {
	return &Metar{IsAuto: false}
}

// Parse parse a metar message.
// Returns an error, if the message could not be parsed.
func (m *Metar) Parse(message string) error {
	m.Message = message

	// remove trend message part,
	// otherwise wind regex will match in the trend also
	noTrendMessage := removeTrendPart(m.Message)

	loc, err := parseLocation(noTrendMessage)
	if err != nil {
		return nil
	}

	m.Location = loc.icao

	wind, err := parseWind(noTrendMessage)
	if err != nil {
		return nil
	}
	m.Wind = wind

	return nil

}

func removeTrendPart(message string) string {
	becmIdx := strings.Index(message, "BCM")
	tempoIdx := strings.Index(message, "TEMPO")
	if becmIdx == -1 && tempoIdx == -1 {
		return message
	}

	if becmIdx != -1 {
		return message[0:becmIdx]
	}
	if tempoIdx != -1 {
		return message[0:tempoIdx]
	}
	return message
}

// HrDateTime return a human readable string of the metar date time field
func (m Metar) HrDateTime() string {
	if m.DateTime == "" || len(m.DateTime) != 7 {
		return ""
	}
	day := m.DateTime[:2]
	time := m.DateTime[2:7]
	return fmt.Sprintf("delivered for the day %v of the month at %v:%v:00(%v)", day, time[0:2], time[2:4], time[4:5])
}
