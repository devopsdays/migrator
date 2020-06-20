package events

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	rice "github.com/GeertJohan/go.rice"
	paths "github.com/devopsdays/devopsdays-cli/helpers/paths"
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

	t.Execute(f, event)
	if err != nil {
		return errors.Wrap(err, "Cannot execute template")
	} else {
		fmt.Println("Created event file for ", event.Name)
	}

	return
}
