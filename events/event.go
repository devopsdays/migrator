package events

import (
	"fmt"

	paths "github.com/devopsdays/migrator/helpers/paths"
)

func GetEvent() {

	thisEvent, _ := GetEventInfo(paths.EventDataPath(paths.GetWebdir(), "chicago", "2019"))

	teamMembers := thisEvent.TeamMembers

	fmt.Println("Team Members")
	for _, person := range teamMembers {
		for key, value := range person {
			fmt.Println(key, ":", value)
		}
		fmt.Println("----------------------")
	}

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

}
