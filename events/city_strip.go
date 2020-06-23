package events

import "strings"

//CityStrip returns a city name with the year removed
func CityStrip(city string) (cityStrip string) { // TODO: change "city" var to eventName or something more clear
	s := strings.SplitN(city, "-", 2)
	cityStrip = s[1]
	return
}

func YearStrip(event string) (yearStrip string) {
	s := strings.Split(event, "-")
	yearStrip = s[0]
	return
}
