package events

import "strings"

//CityStrip returns a city name with the year removed
func CityStrip(city string) (cityStrip string) {
	s := strings.SplitN(city, "-", 2)
	cityStrip = s[1]
	return
}
