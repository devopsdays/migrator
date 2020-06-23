package model

type ContentData struct {
	Title        string   `toml:"title"`
	Description  string   `toml:"description,omitempty"`
	Type         string   `toml:"type,omitempty"`
	Aliases      []string `toml:"aliases,omitempty"`
	SharingImage string   `toml:"sharing_image,omitempty"`
	Speakers     []string `toml:"speakers,omitempty"`
	YouTube      string   `toml:"youtube,omitempty"`
	Vimeo        string   `toml:"vimeo,omitempty"`
	Slideslive   string   `toml:"slideslive,omitempty"`
	Speakerdeck  string   `toml:"speakerdeck,omitempty"`
	Slideshare   string   `toml:"slideshare,omitempty"`
	GoogleSlides string   `toml:"googleslides,omitempty"`
	PDF          string   `toml:"pdf,omitempty"`
	Notist       string   `toml:"notist,omitempty"`
	Slides       string   `toml:"slides,omitempty"`
	Website      string   `toml:"website,omitempty"`
	Twitter      string   `toml:"twitter,omitempty"`
	Facebook     string   `toml:"facebook,omitempty"`
	LinkedIn     string   `toml:"linkedin,omitempty"`
	GitHub       string   `toml:"github,omitempty"`
	GitLab       string   `toml:"gitlab,omitempty"`
	Image        string   `toml:"image,omitempty"`
	Icons        string   `toml:"icons,omitempty"`
	LinkTitle    string   `toml:"linktitle,omitempty"`
	Content      string
}
