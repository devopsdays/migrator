package migrate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func MigrateEvent(city string, year string) {

	thisEvent, _ := GetEventInfo(GetEventDataFilePath(city, year))
	err := CreateEvent(thisEvent)
	if err != nil {
		errors.Wrap(err, "Create event error!")
	}
	MakeOrganizerDirectory(city, year)
	MakeOrganizerIndexFile(city, year)
	MakeOrganizers(thisEvent.TeamMembers)

	// add logic to spin through content directory for the event here and then call CopyPlainEventFile() for it
	CopyPlainEventFile(city, year, "sponsor.md")
	CopyPlainEventFile(city, year, "speakers.md")
	os.MkdirAll(filepath.Join(GetNewEventContentPath(city, year), "speakers"), 0755)
	os.MkdirAll(filepath.Join(GetNewEventContentPath(city, year), "program"), 0755)

	CopyPlainEventFile(city, year, "program/jeff-smith.md")
	CopyPlainEventFile(city, year, "speakers/jeff-smith.md")
	s := []string{year, city}
	slug := strings.Join(s, "-")
	err = MoveStaticDir(slug)
	if err != nil {
		fmt.Println(err)
	}

}
