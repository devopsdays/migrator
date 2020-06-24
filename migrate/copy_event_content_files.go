package migrate

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
	"github.com/pkg/errors"
)

func CopyPlainEventFile(city string, year string, filename string) (err error) {

	sourceContentFilePath := filepath.Join(GetOldEventContentPath(city, year), filename)
	newContentFilePath := filepath.Join(GetNewEventContentPath(city, year), filename)

	thisContent, err := GetContentFileInfo(sourceContentFilePath)
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
	case "event":
		thisContent.Type = "new-event"
	}

	if filename == "contact.md" {
		thisContent.Type = "new-contact"
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
