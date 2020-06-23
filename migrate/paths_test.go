package migrate

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetOldWebDir(t *testing.T) {

	Convey("When checking for the old web directory", t, func() {

		testContentPath := GetOldWebDir()

		Convey("The response should be '/Users/mattstratton/src/github.com/devopsdays/devopsdays-web'", func() {
			So(testContentPath, ShouldEqual, "/Users/mattstratton/src/github.com/devopsdays/devopsdays-web")
		})
	})
}

func TestGetNewWebDir(t *testing.T) {

	Convey("When checking for the new web directory", t, func() {

		testContentPath := GetNewWebDir()

		Convey("The response should be '/Users/mattstratton/src/github.com/devopsdays/devopsdays-web'", func() {
			So(testContentPath, ShouldEqual, "/Users/mattstratton/src/migrate/devopsdays-web")
		})
	})
}

func TestGetOldEventStaticPath(t *testing.T) {

	Convey("Given a city of new-york-city and a year of 2018", t, func() {
		city := "new-york-city"
		year := "2018"

		testStaticPath := GetOldEventStaticPath(city, year)

		Convey("The response should be "+GetOldWebDir()+"/static/events/2018-new-york-city", func() {
			So(testStaticPath, ShouldEqual, GetOldWebDir()+"/static/events/2018-new-york-city")
		})
	})
}

func TestGetNewEventStaticPath(t *testing.T) {

	Convey("Given a city of new-york-city and a year of 2018", t, func() {
		city := "new-york-city"
		year := "2018"

		testStaticPath := GetNewEventStaticPath(city, year)

		Convey("The response should be "+GetNewWebDir()+"/static/events/2018/new-york-city", func() {
			So(testStaticPath, ShouldEqual, GetNewWebDir()+"/static/events/2018/new-york-city")
		})
	})
}

func TestGetOldEventContentPath(t *testing.T) {

	Convey("Given a city of new-york-city and a year of 2018", t, func() {
		city := "new-york-city"
		year := "2018"

		testContentPath := GetOldEventContentPath(city, year)

		Convey("The response should be "+GetOldWebDir()+"/content/events/2018-new-york-city", func() {
			So(testContentPath, ShouldEqual, GetOldWebDir()+"/content/events/2018-new-york-city")
		})
	})
}

func TestGetNewEventContentPath(t *testing.T) {

	Convey("Given a city of new-york-city and a year of 2018", t, func() {
		city := "new-york-city"
		year := "2018"

		testContentPath := GetNewEventContentPath(city, year)

		Convey("The response should be "+GetNewWebDir()+"/content/new-events/2018/new-york-city", func() {
			So(testContentPath, ShouldEqual, GetNewWebDir()+"/content/new-events/2018/new-york-city")
		})
	})
}

func TestGetOldSponsorsPath(t *testing.T) {

	Convey("When checking for the old sponsor directory", t, func() {

		testSponsorPath := GetOldSponsorsPath()

		Convey("The response should be "+GetOldWebDir()+"/data/sponsors", func() {
			So(testSponsorPath, ShouldEqual, GetOldWebDir()+"/data/sponsors")
		})
	})
}

func TestGetNewSponsorsPath(t *testing.T) {

	Convey("When checking for the new sponsor directory", t, func() {

		testSponsorPath := GetNewSponsorsPath()

		Convey("The response should be "+GetNewWebDir()+"/content/sponsors", func() {
			So(testSponsorPath, ShouldEqual, GetNewWebDir()+"/content/sponsors")
		})
	})
}

func TestEventDataFilePath(t *testing.T) {

	Convey("Given a city of New York and a year of 2018", t, func() {
		city := "New York"
		year := "2018"

		testDataFilePath := GetEventDataFilePath(city, year)

		Convey("The response should be "+GetOldWebDir()+"/data/events/2018-new-york.yml", func() {
			So(testDataFilePath, ShouldEqual, GetOldWebDir()+"/data/events/2018-new-york.yml")
		})
	})
}
