package commands

import (
	"github.com/AlecAivazis/survey/v2"
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
			migrateSponsors()
		case "Migrate events":
			migrateEvents()

		}
	}

	return
}
