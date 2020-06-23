package events

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
	"github.com/devopsdays/migrator/helpers"
	"github.com/pkg/errors"
)

func CopyPlainEventFile(city string, year string, filename string) (err error) {
	//TODO: refactor the path variables into something more helpful

	sourceContentFilePath := filepath.Join("/Users/mattstratton/src/github.com/devopsdays/devopsdays-web/content/events/2019-chicago", filename)
	newContentFilePath := filepath.Join("/Users/mattstratton/src/github.com/devopsdays/devopsdays-web/content/new-events/", year, city, filename)

	thisContent, err := helpers.GetContentFileInfo(sourceContentFilePath)
	if err != nil {
		return errors.Wrap(err, "load content failed")
	}

	switch thisContent.Type {
	case "speaker":
		thisContent.Type = "new-speaker"
	case "speakers":
		thisContent.Type = "new-speakers"
	case "talk":
		thisContent.Type = "new-talk"
	}

	templateBox, err := rice.FindBox("templates")
	if err != nil {
		return errors.Wrap(err, "Cannot load content template")
	}

	templateString, err := templateBox.String("content_file.md.tmpl")
	if err != nil {
		return errors.Wrap(err, "Cannot load content template")
	}

	t, err := template.New("content_file.md").Parse(templateString)
	if err != nil {
		return errors.Wrap(err, "Cannot load content template")
	}

	f, err := os.Create(newContentFilePath)
	if err != nil {
		return errors.Wrap(err, "Cannot create content file")
	}

	defer f.Close()

	t.Execute(f, thisContent)
	if err != nil {
		return errors.Wrap(err, "Cannot execute template")
	} else {
		fmt.Println("Created content file ", newContentFilePath)
	}

	return nil

}
