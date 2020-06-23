package migrate

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/pkg/errors"
)

// func MigrateOrganizers(city string, year string, organizers []map[string]string) (err errors) {

// 	return nil
// }

func MakeOrganizerDirectory(city string, year string) (err error) {
	err = os.MkdirAll(filepath.Join((GetNewEventContentPath(city, year)), "organizers"), 0755)
	if err != nil {
		return errors.Wrap(err, "Make event organizer directory failed")
	}
	return nil
}

func MakeOrganizerIndexFile(city string, year string) (err error) {
	// find a rice.Box
	// to compile, cd to organizers directory and run `rice embed-go`
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		return errors.Wrap(err, "content template find failed")
	}
	// get file contents as string
	templateString, err := templateBox.String("headless.index.md.tmpl")
	if err != nil {
		return errors.Wrap(err, "Cannot load organizer index template")
	}

	t, err := template.New("headless.index.md").Parse(templateString)
	if err != nil {
		return errors.Wrap(err, "Cannot load organizer template")
	}

	f, err := os.Create(filepath.Join((GetNewEventContentPath(city, year)), "organizers", "index.md"))
	if err != nil {
		return errors.Wrap(err, "Cannot create organizer index file")
	}

	defer f.Close()

	t.Execute(f, "")
	if err != nil {
		return errors.Wrap(err, "Cannot execute template")
	} else {
		fmt.Println("Created organizer index file for ", year, " ", city)
	}

	return
}

func MakeOrganizers(organizers []map[string]string) (err error) {

	for _, person := range organizers {
		var thisOrganizer Organizer
		for key, value := range person {
			switch key {
			case "name":
				thisOrganizer.Title = value
			case "twitter":
				thisOrganizer.Twitter = value
			case "linkedin":
				thisOrganizer.LinkedIn = value
			case "github":
				thisOrganizer.GitHub = value
			case "website":
				thisOrganizer.Website = value
			case "employer":
				thisOrganizer.Employer = value
			case "role":
				thisOrganizer.Role = value
			case "image":
				thisOrganizer.Image = value
			case "gitlab":
				thisOrganizer.GitLab = value
			case "facebook":
				thisOrganizer.Facebook = value
			case "bio":
				thisOrganizer.Bio = value
			}
		}
		MakeOrganizerFile("chicago", "2019", thisOrganizer)

	}
	return
}

func MakeOrganizerFile(city string, year string, organizer Organizer) (err error) {

	// find a rice.Box
	// to compile, cd to organizers directory and run `rice embed-go`
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		return errors.Wrap(err, "content template find failed")
	}
	// get file contents as string
	templateString, err := templateBox.String("organizer.md.tmpl")
	if err != nil {
		return errors.Wrap(err, "Cannot load organizer organizer template")
	}

	t, err := template.New("organizer.md").Parse(templateString)
	if err != nil {
		return errors.Wrap(err, "Cannot load organizer template")
	}

	organizerFileName := strings.ReplaceAll((strings.ToLower(organizer.Title)), " ", "-")
	organizerFileName += ".md"
	f, err := os.Create(filepath.Join((GetNewEventContentPath(city, year)), "organizers", organizerFileName))
	if err != nil {
		return errors.Wrap(err, "Cannot create event file")
	}

	defer f.Close()

	t.Execute(f, organizer)
	if err != nil {
		return errors.Wrap(err, "Cannot execute template")
	} else {
		fmt.Println("Created organizer file for ", year, " ", city, "-", organizer.Title)
	}

	return

}

type Organizer struct {
	Title    string
	Twitter  string
	LinkedIn string
	GitHub   string
	Website  string
	Employer string
	Role     string
	Image    string
	GitLab   string
	Facebook string
	Bio      string
}
