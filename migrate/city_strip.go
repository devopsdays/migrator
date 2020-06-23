package migrate

import "strings"

// CityStrip returns an event name with the year removed
func CityStrip(event string) (cityStrip string) {
	s := strings.SplitN(event, "-", 2)
	cityStrip = s[1]
	return
}

func YearStrip(event string) (yearStrip string) {
	s := strings.Split(event, "-")
	yearStrip = s[0]
	return
}
