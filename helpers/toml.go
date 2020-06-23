package helpers

import (
	"io/ioutil"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/devopsdays/migrator/model"
	"github.com/gernest/front"
	"github.com/pkg/errors"
)

// TOMLHandler decodes TOML string from a content file into a go map[string]interface{}
func TOMLHandler(front string) (map[string]interface{}, error) {

	var thisContent model.ContentData
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

func GetContentFileInfo(file string) (contentData *model.ContentData, err error) {

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

	contentData = &model.ContentData{
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
