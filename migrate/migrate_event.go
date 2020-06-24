package migrate

import (
	"fmt"
	"io/ioutil"
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

	if EventHasWelcomeFile(city, year) {
		MakeOrganizerDirectory(city, year)
		MakeOrganizerIndexFile(city, year)
		MakeOrganizers(thisEvent.TeamMembers)

		if _, err := os.Stat(filepath.Join(GetOldEventContentPath(city, year), "speakers")); err == nil {

			os.MkdirAll(filepath.Join(GetNewEventContentPath(city, year), "speakers"), 0755)
			speakerFiles, _ := ioutil.ReadDir(filepath.Join(GetOldEventContentPath(city, year), "speakers"))
			for _, f := range speakerFiles {
				CopyPlainEventFile(city, year, filepath.Join("speakers", f.Name()))
			}
		}
		if _, err := os.Stat(filepath.Join(GetOldEventContentPath(city, year), "speakers")); err == nil {
			os.MkdirAll(filepath.Join(GetNewEventContentPath(city, year), "program"), 0755)
			programFiles, _ := ioutil.ReadDir(filepath.Join(GetOldEventContentPath(city, year), "program"))
			for _, f := range programFiles {
				CopyPlainEventFile(city, year, filepath.Join("program", f.Name()))
			}
		}

		eventFiles, _ := ioutil.ReadDir(filepath.Join(GetOldEventContentPath(city, year)))
		for _, f := range eventFiles {
			if f.IsDir() {
				fmt.Println("directory!")
			} else {
				if f.Name() != "welcome.md" {
					CopyPlainEventFile(city, year, f.Name())

				}
			}
		}

		// delete content/events on destination
		fmt.Println("removing ", GetLegacyEventContentPath(city, year))
		os.RemoveAll(GetLegacyEventContentPath(city, year))
		dataFileName := thisEvent.Name + ".yml"
		dataFilePath := filepath.Join(GetNewWebDir(), "data", "events", dataFileName)
		os.Remove(dataFilePath)

	}

	s := []string{year, city}
	slug := strings.Join(s, "-")
	err = MoveStaticDir(slug)
	if err != nil {
		fmt.Println(err)
	}

}
