package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
	"gopkg.in/yaml.v2"
)

type Sponsor struct {
	Name    string `yaml:"name"`
	Website string `yaml:"url"`
	Twitter string `yaml:"twitter,omitempty"`
}

func main() {
	fmt.Printf("It's the migrator")
}

func convertSponsors(sourceDir, destDir string) (err error) {

	// find the rice.Box
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		log.Fatal(err)
	}

	// get file contents as a string
	templateString, err := templateBox.String("sponsor.md.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("guest.md").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}

	files, _ := ioutil.ReadDir(sourceDir)
	for _, f := range files {

		var sponsor Sponsor
		sponsorFile := filepath.Join(sourceDir, f.Name())
		yamlFile, err := ioutil.ReadFile(sponsorFile)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(yamlFile, &sponsor)

	}

	return
}
