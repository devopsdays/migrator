package migrate

import (
	"fmt"
	"html"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
	"github.com/pkg/errors"
)

type EventData struct {
	Name                  string              `yaml:"name"`
	Year                  string              `yaml:"year"`
	City                  string              `yaml:"city"`
	EventGroup            string              `yaml:"event_group,omitempty"`
	EventTwitter          string              `yaml:"event_twitter"`
	EventDescription      string              `yaml:"description,omitempty"`
	GoogleAnalytics       string              `yaml:"ga_tracking_id,omitempty"`
	SpeakersVerbose       string              `yaml:"speakers_verbose,omitempty"`
	Cancel                string              `yaml:"cancel,omitempty"`
	StartDate             string              `yaml:"startdate,omitempty"`
	EndDate               string              `yaml:"enddate,omitempty"`
	CFPDateStart          string              `yaml:"cfp_date_start,omitempty"`
	CFPDateEnd            string              `yaml:"cfp_date_end,omitempty"`
	CFPDateAnnounce       string              `yaml:"cfp_data_announce,omitempty"`
	CFPOpen               string              `yaml:"cfp_open,omitempty"`
	CFPLink               string              `yaml:"cfp_link,omitempty"`
	RegistrationDateStart string              `yaml:"registration_date_start,omitempty"`
	RegistrationDateEnd   string              `yaml:"registration_date_end,omitempty"`
	RegistrationClosed    string              `yaml:"registration_closed,omitempty"`
	RegistrationLink      string              `yaml:"registration_link,omitempty"`
	SponsorLink           string              `yaml:"sponsor_link,omitempty"`
	MastheadBackground    string              `yaml:"masthead_background,omitempty"`
	SharingImage          string              `yaml:"sharing_image,omitempty"`
	Coordinates           string              `yaml:"coordinates,omitempty"`
	Location              string              `yaml:"location,omitempty"`
	LocationAddress       string              `yaml:"location_address,omitempty"`
	NavElements           []map[string]string `yaml:"nav_elements"`
	TeamMembers           []map[string]string `yaml:"team_members"`
	OrganizerEmail        string              `yaml:"organizer_email"`
	ProposalEmail         string              `yaml:"proposal_email"`
	Sponsors              []map[string]string `yaml:"sponsors"`
	SponsorsAccepted      string              `yaml:"sponsors_accepted"`
	SponsorLevels         []map[string]string `yaml:"sponsor_levels"`
	Content               string
}

func CreateEvent(event EventData) (err error) {

	// find a rice.Box
	// to compile, cd to events directory and run `rice embed-go`
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		return errors.Wrap(err, "content template find failed")
	}
	// get file contents as string
	templateString, err := templateBox.String("event._index.md.tmpl")
	if err != nil {
		return errors.Wrap(err, "Cannot load event template")
	}

	t, err := template.New("event._index.md").Parse(templateString)
	if err != nil {
		return errors.Wrap(err, "Cannot load event template")
	}

	fmt.Println("Event name is ", event.Name)

	eventContentPath := GetNewEventContentPath(CityStrip(event.Name), event.Year)
	err = os.MkdirAll(eventContentPath, 0755)
	if err != nil {
		return errors.Wrap(err, "make event content directory failed")
	}

	newEventFilePath := filepath.Join(eventContentPath, "_index.md")

	f, err := os.Create(newEventFilePath)
	if err != nil {
		return errors.Wrap(err, "Cannot create event file")
	}

	defer f.Close()

	if EventHasWelcomeFile(CityStrip(event.Name), event.Year) {

		sourceContentFilePath := filepath.Join(GetOldEventContentPath(CityStrip(event.Name), event.Year), "welcome.md")

		fmt.Println("source welcome file is ", sourceContentFilePath)

		thisContent, err := GetContentFileInfo(sourceContentFilePath)
		if err != nil {
			return errors.Wrap(err, "load content failed")
		}
		city_slug := event.Name
		event_year := event.Year
		event_city := CityStrip(city_slug)
		new_path := fmt.Sprintf("%s/%s", event_year, event_city)
		// fmt.Println("event slug is: ", city_slug)
		// fmt.Println("New path is: ", new_path)
		event.Content = thisContent.Content
		event.Content = strings.ReplaceAll(event.Content, city_slug, new_path)
		event.Content = html.UnescapeString(event.Content)
	}

	t.Execute(f, event)
	if err != nil {
		return errors.Wrap(err, "Cannot execute template")
	} else {
		fmt.Println("Created event file for ", event.Name)
		dataFileName := event.Name + ".yml"
		dataFilePath := filepath.Join(GetNewWebDir(), "data", "events", dataFileName)
		// fmt.Println("Deleting file at ", GetEventDataFilePath(CityStrip(event.Name), event.Year))
		os.Remove(dataFilePath)

	}

	return

}

func EventHasWelcomeFile(city string, year string) (hasWelcomeFile bool) {
	if _, err := os.Stat(filepath.Join(GetOldEventContentPath(city, year), "welcome.md")); err == nil {
		fmt.Println("welcome file exists")
		return true
	}
	return false
}
