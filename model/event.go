package model

type EventData struct {
	Name                  string              `yaml:"name"`
	Year                  string              `yaml:"year"`
	City                  string              `yaml:"city"`
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
}
