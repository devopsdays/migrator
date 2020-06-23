package migrate

import (
	"io/ioutil"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/gernest/front"
	"github.com/pkg/errors"
)

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

// TOMLHandler decodes TOML string from a content file into a go map[string]interface{}
func TOMLHandler(front string) (map[string]interface{}, error) {

	var thisContent ContentData
	if _, err := toml.Decode(front, &thisContent); err != nil {
		return nil, errors.Wrap(err, "TOML decoding failed")
	}

	x := map[string]interface{}{
		"Title":        thisContent.Title,
		"Description":  thisContent.Description,
		"Type":         thisContent.Type,
		"Aliases":      thisContent.Aliases,
		"SharingImage": thisContent.SharingImage,
		"Speakers":     thisContent.Speakers,
		"YouTube":      thisContent.YouTube,
		"Vimeo":        thisContent.Vimeo,
		"Slideslive":   thisContent.Slideslive,
		"Slideshare":   thisContent.Slideshare,
		"Speakerdeck":  thisContent.Speakerdeck,
		"GoogleSlides": thisContent.GoogleSlides,
		"PDF":          thisContent.PDF,
		"Notist":       thisContent.Notist,
		"Slides":       thisContent.Slides,
		"Website":      thisContent.Website,
		"Twitter":      thisContent.Twitter,
		"Facebook":     thisContent.Facebook,
		"LinkedIn":     thisContent.LinkedIn,
		"GitHub":       thisContent.GitHub,
		"GitLab":       thisContent.GitLab,
		"Image":        thisContent.Image,
		"Icons":        thisContent.Icons,
		"LinkTitle":    thisContent.LinkTitle,
	}
	return x, nil
}

func GetContentFileInfo(file string) (contentData *ContentData, err error) {

	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.Wrap(err, "load file failed")
	}

	m := front.NewMatter()
	m.Handle("+++", TOMLHandler)

	if err != nil {
		return nil, errors.Wrap(err, "TOML handler failed")
	}

	f, body, err := m.Parse(strings.NewReader(string(dat)))
	if err != nil {
		return nil, errors.Wrap(err, "Get content info failed")
	}

	contentData = &ContentData{
		Title:        f["Title"].(string),
		Description:  f["Description"].(string),
		Type:         f["Type"].(string),
		Aliases:      f["Aliases"].([]string),
		SharingImage: f["SharingImage"].(string),
		Speakers:     f["Speakers"].([]string),
		YouTube:      f["YouTube"].(string),
		Vimeo:        f["Vimeo"].(string),
		Slideslive:   f["Slideslive"].(string),
		Speakerdeck:  f["Speakerdeck"].(string),
		Slideshare:   f["Slideshare"].(string),
		GoogleSlides: f["GoogleSlides"].(string),
		PDF:          f["PDF"].(string),
		Notist:       f["Notist"].(string),
		Slides:       f["Slides"].(string),
		Website:      f["Website"].(string),
		Twitter:      f["Twitter"].(string),
		Facebook:     f["Facebook"].(string),
		LinkedIn:     f["LinkedIn"].(string),
		GitHub:       f["GitHub"].(string),
		GitLab:       f["GitLab"].(string),
		Image:        f["Image"].(string),
		Icons:        f["Icons"].(string),
		LinkTitle:    f["LinkTitle"].(string),
		Content:      body,
	}

	return contentData, nil
}
