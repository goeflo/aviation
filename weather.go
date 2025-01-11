package aviation

//import (
//	"fmt"
//	"regexp"
//)
//
//var weatherPattern = "(?P<intensity>\\+|-)?(?<weather>(BZ|BL|BR|DR|DS|DU|DZ|FG|FC|FU|FZ|GR|GS|HZ|IC|MI|PL|PO|RA|SA|SG|SH|SN|SQ|SS|TS|VA|VC|UP|RE){1,3})"
//
//type Weather struct {
//	intensity  string
//	indication []string
//}
//
//func parseWeather(message string) (*Weather, error) {
//	weatherRegex := regexp.MustCompile(weatherPattern)
//
//	if !weatherRegex.MatchString(message) {
//		return nil, fmt.Errorf("no weather information found")
//	}
//	fmt.Printf("found weather %v\n", message)
//	weather := &Weather{}
//	//weatherInformation := weatherRegex.FindAllStringSubmatch(message, -1)[0][1]
//	weatherInformation := weatherRegex.FindStringSubmatch(message)
//	fmt.Printf("weather information: %s\n", weatherInformation)
//
//	if weatherInformation[0] == "+" || weatherInformation[0] == "-" {
//		weather.intensity = weatherInformation[0]
//		weatherInformation = weatherInformation[1:]
//	}
//
//	indications := len(weatherInformation) / 2
//	for i := 0; i < indications; i += 2 {
//		weather.indication = append(weather.indication, "weatherInformation")
//	}
//	return weather, nil
//
//}
//
