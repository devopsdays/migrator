package events

import (
	"fmt"

	paths "github.com/devopsdays/migrator/helpers/paths"
	"github.com/devopsdays/migrator/organizers"
)

func GetEvent(city string, year string) {

	thisEvent, _ := GetEventInfo(paths.EventDataPath(paths.GetWebdir(), city, year))

	teamMembers := thisEvent.TeamMembers

	navElements := thisEvent.NavElements
	fmt.Println("Navigation Elements")
	for _, element := range navElements {
		for key, value := range element {
			fmt.Println(key, ":", value)
			if key == "url" {
				fmt.Println("there is a url")
			}
		}
		fmt.Println("----------------------")
	}

	CreateEvent(thisEvent)
	organizers.MakeOrganizers(teamMembers)

}
