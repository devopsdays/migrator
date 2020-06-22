package events

import "github.com/devopsdays/migrator/organizers"

func MigrateEvent(city string, year string) {
	organizers.MakeOrganizerDirectory(city, year)
	organizers.MakeOrganizerIndexFile(city, year)
	GetEvent(city, year)
}
