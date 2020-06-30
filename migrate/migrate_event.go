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
		fmt.Println("hello i am about to create the organizer directory")
		MakeOrganizerDirectory(city, year)
		MakeOrganizerIndexFile(city, year)
		MakeOrganizers(thisEvent.TeamMembers, city, year)
		fmt.Println(thisEvent.TeamMembers)

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

		// check to see if there is a data/speakers/yyyy/city directory. If so...
		fmt.Println("Speaker data files directory is ", GetEventSpeakerDataFilesPath(city, year))
		if _, err := os.Stat(GetEventSpeakerDataFilesPath(city, year)); err == nil {
			newSpeakerDir := filepath.Join(GetNewEventContentPath(city, year), "speakers")
			oldProgramDir := filepath.Join(GetOldEventContentPath(city, year), "program")
			newProgramDir := filepath.Join(GetNewEventContentPath(city, year), "program")
			fmt.Println("Old speaker directory is ", GetEventSpeakerDataFilesPath(city, year))
			fmt.Println("New speaker directory is", newSpeakerDir)
			fmt.Println("Old Program directory is ", oldProgramDir)
			fmt.Println("New Program directory is ", newProgramDir)
			os.MkdirAll(newSpeakerDir, 0755)
			os.MkdirAll(newProgramDir, 0755)

			ConvertSpeakers(GetEventSpeakerDataFilesPath(city, year), newSpeakerDir, city, year)
			ConvertTalks(oldProgramDir, newProgramDir)

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
