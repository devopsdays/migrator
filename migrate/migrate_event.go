package migrate

import (
	"fmt"
	"strings"
)

func MigrateEvent(city string, year string) {
	MakeOrganizerDirectory(city, year)
	MakeOrganizerIndexFile(city, year)
	GetEvent(city, year)
	// add thisEvent := GetEvent(city, year) here
	// add CreateEvent(thisEvent) here
	// add MakeOrganizers(teamMembers) here - or just call the MigrateOrganizers(city,year) function
	// add logic to spin through content directory for the event here and then call CopyPlainEventFile() for it
	CopyPlainEventFile(city, year, "sponsor.md")
	CopyPlainEventFile(city, year, "speakers.md")
	CopyPlainEventFile(city, year, "program/jeff-smith.md")
	CopyPlainEventFile(city, year, "speakers/jeff-smith.md")
	s := []string{year, city}
	slug := strings.Join(s, "-")
	err := MoveStaticDir(slug)
	if err != nil {
		fmt.Println(err)
	}

}
