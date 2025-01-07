package aviation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetar(t *testing.T) {

	tests := map[string]struct {
		msg      string
		err      error
		icao     string
		dateTime string
	}{
		"LBBG": {
			msg:      "METAR LBBG 041600Z 12012MPS 090V150 1400 R04/P1500N R22/P1500U +SN BKN022 OVC050 M04/M07 Q1020 NOSIG 8849//91=",
			err:      nil,
			icao:     "LBBG",
			dateTime: "041600Z",
		},
		"EDDB": {
			msg:      "METAR EDDB 070920Z AUTO 22016KT 9999 FEW036 BKN043 08/01 Q0998 TEMPO 20020G35KT=",
			err:      nil,
			icao:     "EDDB",
			dateTime: "070920Z",
		},
		"EDDW I": {
			msg: "EDDW 160650Z 00000KT 0100 R14/0250N R16/0250V0400U R28/0300D FZFG VV001 M01/M01 Q1012 BECMG 0600 BKN005",
		},
		"EDDW II": {
			msg: "EDDW 231820Z 22010G22KT 190V260 8000 -TSRA FEW010 SCT025CB BKN050 23/15 Q1010 WS ALL RWY BECMG FM1900  NSW",
		},
	}

	metar := NewMetar()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := metar.Parse(test.msg)
			assert.Equal(t, test.err, err, name)
			assert.Equal(t, test.icao, metar.Icao, name)

			fmt.Printf("%v\n", metar.HrDateTime())
		})
	}

}