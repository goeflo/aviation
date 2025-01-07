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
*/
type Metar struct {
	Message  string
	Icao     string
	DateTime string
	IsAuto   bool
}

// new metar struct
func NewMetar() *Metar {
	return &Metar{IsAuto: false}
}

// Parse parse a metar message.
// Returns an error, if the message could not be parsed.
func (m *Metar) Parse(newMessage string) error {
	m.Message = newMessage

	messageTokens := strings.Split(m.Message, " ")

	token := 0

	m.Icao = messageTokens[token]
	if len(m.Icao) != 4 {
		return fmt.Errorf("icao code '%v' does not have the required length of 4 char", m.Icao)
	}

	token++
	m.DateTime = messageTokens[token]

	token++
	if messageTokens[token] == "AUTO" {
		m.IsAuto = true
	}

	return nil

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
