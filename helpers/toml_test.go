package helpers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTOMLHandler(t *testing.T) {

	Convey("Given a set of frontmatter", t, func() {
		var text = `
		Title = "This is my title"
		Type = "welcome"
		aliases = ["/events/2019-chicago/","something/else"]
		Description = "This is the description."
		image = "speaker-image.jpg"
		Twitter = "mattstratton"
		linktitle = "matt-stratton"
		`

		testContentData, _ := TOMLHandler(text)

		for key, value := range testContentData {
			switch key {
			case "Title":
				Convey("The response should be 'This is my title'", func() {
					So(value, ShouldEqual, "This is my title")

				})
			case "Type":
				Convey("The response should be 'welcome'", func() {
					So(value, ShouldEqual, "welcome")

				})

			case "Image":
				Convey("The response should be 'speaker-image.jpg'", func() {
					So(value, ShouldEqual, "speaker-image.jpg")

				})
			}
		}

	})
}
