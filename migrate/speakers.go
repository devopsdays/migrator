package migrate

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// this will read in a speaker yaml data file and create the associated page in content/events/YYYY/city/speakers/ directory

// note that this should be called in the larger event migration; but only if there is a `data/speakers/yyyy/city` directory for that year/city.

// if this gets run, it should also do a conversion for any associated program page; need to think about the race condition of what happens when.
// The program files will *not* have been moved, so this should run against the old content path (this is what it does in the function below)

// note! speaker data files do not have image. Should check static/events/yyyy-city/speakers/SPEAKERNAME.jpg" and if it exists, set image = "speakername.jpg"

// also note - these functions assume that the destination directories have been created and that the source directories exist, so no need
// to check that!

type Speaker struct {
	Name    string `yaml:"name"`
	Twitter string `yaml:"twitter,omitempty"`
	Bio     string `yaml:"bio,omitempty"`
	Image   string `yaml:"image,omitempty"`
}

func ConvertSpeakers(sourceDir, destDir, city, year string) (err error) {
	// this should take in city and year not directories?

	// find the rice.Box
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		log.Fatal(err)
	}

	// get file contents as a string
	templateString, err := templateBox.String("speaker.md.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("speaker.md").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}

	files, _ := ioutil.ReadDir(sourceDir)

	for _, f := range files {

		var speaker Speaker
		speakerFile := filepath.Join(sourceDir, f.Name())
		speakerSlug := strings.TrimSuffix(f.Name(), path.Ext(f.Name()))
		yamlFile, err := ioutil.ReadFile(speakerFile)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(yamlFile, &speaker)
		if err != nil {
			panic(err)
		}
		fmt.Println(speaker.Name)
		s := []string{speakerSlug, ".md"}
		newSpeakerFilePath := filepath.Join(destDir, strings.Join(s, ""))
		fmt.Println(newSpeakerFilePath)

		// data := struct {
		// 	Name    string
		// 	Twitter string
		// 	Bio     string
		// }{
		// 	speaker.Name,
		// 	speaker.Twitter,
		// 	speaker.Bio,
		// }
		f, err := os.Create(newSpeakerFilePath)
		if err != nil {
			fmt.Println(err)
		}
		s = []string{speakerSlug, ".jpg"}
		speakerImagePath := filepath.Join(GetOldEventStaticPath(city, year), "speakers", strings.Join(s, ""))
		fmt.Println("Speaker image path is ", speakerImagePath)
		if _, err := os.Stat(speakerImagePath); err == nil {
			fmt.Println("I should set the image!")
			speaker.Image = strings.Join(s, "")
			fmt.Println(speaker)
		}

		t.Execute(f, speaker)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Created speaker file for ", speaker.Name)
		}
		f.Close()
	}

	return

}

func ConvertTalks(sourceDir, destDir string) (err error) {

	// find the rice.Box
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		log.Fatal(err)
	}

	// get file contents as a string
	templateString, err := templateBox.String("talk.md.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("talk.md").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}

	files, _ := ioutil.ReadDir(sourceDir)

	for _, f := range files {

		fmt.Println("Source talk file is ", f.Name())

		sourceContentFilePath := filepath.Join(sourceDir, f.Name())
		newContentFilePath := filepath.Join(destDir, f.Name())

		speakerSlug := strings.TrimSuffix(f.Name(), path.Ext(f.Name()))

		thisContent, err := GetContentFileInfo(sourceContentFilePath)
		if err != nil {
			return errors.Wrap(err, "load content failed")
		}

		f, err := os.Create(newContentFilePath)
		if err != nil {
			return errors.Wrap(err, "Cannot create content file")
		}

		thisContent.Speakers = append(thisContent.Speakers, speakerSlug)

		t.Execute(f, thisContent)
		if err != nil {
			return errors.Wrap(err, "Cannot execute template")
		} else {
			fmt.Println("Created content file ", newContentFilePath)
		}
		f.Close()

	}

	return
}
