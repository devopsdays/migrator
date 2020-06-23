package events

import (
	"fmt"
	"html"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
	"github.com/devopsdays/migrator/helpers"
	paths "github.com/devopsdays/migrator/helpers/paths"
	"github.com/devopsdays/migrator/model"
	"github.com/pkg/errors"
)

func CreateEvent(event model.EventData) (err error) {

	// find a rice.Box
	// to compile, cd to events directory and run `rice embed-go`
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		return errors.Wrap(err, "content template find failed")
	}
	// get file contents as string
	templateString, err := templateBox.String("_index.md.tmpl")
	if err != nil {
		return errors.Wrap(err, "Cannot load event template")
	}

	t, err := template.New("_index.md").Parse(templateString)
	if err != nil {
		return errors.Wrap(err, "Cannot load event template")
	}

	eventContentPath := filepath.Join(paths.GetWebdir(), "content", "new-events", event.Year, CityStrip(event.Name))
	err = os.MkdirAll(eventContentPath, 0755)
	if err != nil {
		return errors.Wrap(err, "make event content directory failed")
	}

	newEventFilePath := filepath.Join(eventContentPath, "_index.md")

	f, err := os.Create(newEventFilePath)
	if err != nil {
		return errors.Wrap(err, "Cannot create event file")
	}

	defer f.Close()

	// TODO: Add a check for the welcome.md file and then do the stuff. Might be better to move this into another function

	sourceContentFilePath := filepath.Join(paths.GetWebdir(), "content", "events", event.Name, "welcome.md")

	thisContent, err := helpers.GetContentFileInfo(sourceContentFilePath)
	if err != nil {
		return errors.Wrap(err, "load content failed")
	}

	city_slug := event.Name
	event_year := event.Year
	event_city := CityStrip(city_slug)
	new_path := fmt.Sprintf("%s/%s", event_year, event_city)
	// fmt.Println("event slug is: ", city_slug)
	// fmt.Println("New path is: ", new_path)
	event.Content = thisContent.Content
	event.Content = strings.ReplaceAll(event.Content, city_slug, new_path)
	event.Content = html.UnescapeString(event.Content)

	t.Execute(f, event)
	if err != nil {
		return errors.Wrap(err, "Cannot execute template")
	} else {
		fmt.Println("Created event file for ", event.Name)
	}

	return
}
