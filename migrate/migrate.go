// Package migrate performs all the migration tasks for new devopsdays theme
package migrate

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

func MigrateAllEvents() {
	// spin through all event data files

	// call migrate event from what we know

	files, _ := ioutil.ReadDir(GetOldEventDataFilesPath())

	for _, f := range files {
		eventSlug := strings.TrimSuffix(f.Name(), path.Ext(f.Name()))
		fmt.Println(eventSlug)

		MigrateEvent(CityStrip(eventSlug), YearStrip(eventSlug))

	}

}
