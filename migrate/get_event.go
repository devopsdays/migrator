package migrate

import (
	"fmt"
)

func GetEvent(city string, year string) { //TODO: this should return an event struct/object, and the CreateEvent() call should be moved to MigrateEvent()

	thisEvent, _ := GetEventInfo(GetEventDataFilePath(city, year))

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
	MakeOrganizers(teamMembers) // TODO: move this to MigrateEvent() (remember to grab setting of teamMembers too)

}
