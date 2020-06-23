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

// run rice embed-go to compile the templates
// to compile, cd to migrate directory and run `rice embed-go`

type Sponsor struct {
	Name    string `yaml:"name"`
	Website string `yaml:"url"`
	Twitter string `yaml:"twitter,omitempty"`
}

func PrepSponsorDir(sponsorDir string) (err error) {

	// create the sponsor directory

	err = os.MkdirAll(sponsorDir, 0755)

	if err != nil {
		errors.Wrap(err, "Make sponsor content directory failed")
	}

	// find a rice.Box
	// to compile, cd to organizers directory and run `rice embed-go`
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		return errors.Wrap(err, "content template find failed")
	}
	// get file contents as string
	templateString, err := templateBox.String("headless.index.md.tmpl")
	if err != nil {
		return errors.Wrap(err, "Cannot load sponsor index template")
	}

	t, err := template.New("headless.index.md").Parse(templateString)
	if err != nil {
		return errors.Wrap(err, "Cannot load sponsor template")
	}

	f, err := os.Create(filepath.Join(sponsorDir, "index.md"))
	if err != nil {
		return errors.Wrap(err, "Cannot create sponsor index file")
	}

	defer f.Close()

	t.Execute(f, "")
	if err != nil {
		return errors.Wrap(err, "Cannot execute template")
	} else {
		fmt.Println("Created sponsor index file")
	}

	return nil

}
func ConvertSponsors(sourceDir, destDir string) (err error) {

	PrepSponsorDir(destDir)

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
