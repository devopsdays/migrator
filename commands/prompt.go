package commands

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/devopsdays/migrator/events"
	"github.com/devopsdays/migrator/sponsors"
)

func mainPrompt() (err error) {
	var selection string

	for selection != "Quit the application" {
		prompt := &survey.Select{
			Message: "Select an action:",
			Options: []string{
				"Migrate sponsors",
				"Migrate events",
				"Quit the application",
			},
		}
		err := survey.AskOne(prompt, &selection, nil)
		if err != nil {
			break
		}
		switch selection {
		case "Migrate sponsors":
			fmt.Println("migrate sponsors")
			sponsors.ConvertSponsors("/Users/mattstratton/src/github.com/devopsdays/devopsdays-web/data/sponsors/", "/Users/mattstratton/src/github.com/devopsdays/devopsdays-web/content/new-sponsors/")
		case "Migrate events":
			events.MigrateEvent("chicago", "2019")

		}
	}

	return
}
