package events

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckEvent(t *testing.T) {

	Convey("Given an event city and year", t, func() {
		city := "Chicago"
		year := "2019"
		Convey("When the event exists", func() {
			result := CheckEvent(city, year)
			Convey("Then the result should be true", func() {
				So(result, ShouldEqual, true)
			})
		})
	})

	Convey("Given an event city and year", t, func() {
		city := "Canterlot"
		year := "2016"
		Convey("When the event does not exist", func() {
			result := CheckEvent(city, year)
			Convey("Then the result should be false", func() {
				So(result, ShouldEqual, false)
			})
		})
	})
}
