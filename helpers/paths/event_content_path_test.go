package paths

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEventContentPath(t *testing.T) {

	Convey("Given a city of New York City and a year of 2018", t, func() {
		city := "New York City"
		year := "2018"

		testContentPath := EventContentPath(city, year)

		Convey("The response should be "+GetWebdir()+"/content/events/2018-new-york-city", func() {
			So(testContentPath, ShouldEqual, GetWebdir()+"/content/events/2018-new-york-city")
		})
	})
}

func TestEventNewContentPath(t *testing.T) {

	Convey("Given a city of New York City and a year of 2018", t, func() {
		city := "new-york-city"
		year := "2018"

		testContentPath := EventNewContentPath(city, year)

		Convey("The response should be "+GetWebdir()+"/content/new-events/2018/new-york-city", func() {
			So(testContentPath, ShouldEqual, GetWebdir()+"/content/new-events/2018/new-york-city")
		})
	})
}
