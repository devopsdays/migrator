package sponsors

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
	"gopkg.in/yaml.v2"
)

// run rice embed-go to compile the templates
// to compile, cd to sponsors directory and run `rice embed-go`

type Sponsor struct {
	Name    string `yaml:"name"`
	Website string `yaml:"url"`
	Twitter string `yaml:"twitter,omitempty"`
}

func ConvertSponsors(sourceDir, destDir string) (err error) {

	// find the rice.Box
	templateBox, err := rice.FindBox("sponsor-templates")
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
		// TODO just strip the extension, don't check for it, because it could be yml or yaml
		sponsorSlug := strings.TrimSuffix(f.Name(), path.Ext(f.Name()))
		yamlFile, err := ioutil.ReadFile(sponsorFile)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(yamlFile, &sponsor)
		if err != nil {
			panic(err)
		}
		fmt.Println(sponsor.Name)
		s := []string{sponsorSlug, ".md"}
		newSponsorFilePath := filepath.Join(destDir, strings.Join(s, ""))
		fmt.Println(newSponsorFilePath)

		data := struct {
			Name    string
			Website string
			Twitter string
		}{
			sponsor.Name,
			sponsor.Website,
			sponsor.Twitter,
		}
		f, err := os.Create(newSponsorFilePath)
		if err != nil {
			fmt.Println(err)
		}

		t.Execute(f, data)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Created sponsor file for ", sponsor.Name)
		}
		f.Close()

	}

	return
}
