package migrate

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCityStrip(t *testing.T) {

	Convey("Given the slug 2019-new-york-city", t, func() {
		event := "2019-new-york-city"
		Convey("The result should be new-york-city", func() {
			So(CityStrip(event), ShouldEqual, "new-york-city")
		})

	})
}

func TestYearStrip(t *testing.T) {

	Convey("Given the slug 2019-new-york-city", t, func() {
		event := "2019-new-york-city"
		Convey("The result should be 2019", func() {
			So(YearStrip(event), ShouldEqual, "2019")
		})

	})
}

// add a thing
