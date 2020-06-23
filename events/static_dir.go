package events

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/devopsdays/migrator/helpers/paths"
	"github.com/pkg/errors"
)

func MoveStaticDir(slug string) (err error) {
	year := YearStrip(slug)
	city := CityStrip(slug)
	oldStaticDir := filepath.Join(paths.GetStaticDir(), "events", slug)
	oldStaticDir += "/"
	newStaticDir := filepath.Join(paths.GetStaticDir(), "events", year, city)
	newStaticDir += "/"
	if _, err := os.Stat(newStaticDir); err == nil {
		fmt.Println("New directory already exists")
		return nil
	}
	err = os.MkdirAll(filepath.Join(paths.GetStaticDir(), "events", year), 0755)
	if err != nil {
		return errors.Wrap(err, "make year static directory failed")
	}
	fmt.Println("Created new static directory at ", newStaticDir)
	err = os.Rename(oldStaticDir, newStaticDir)
	if err != nil {
		return errors.Wrap(err, "rename directory failed")
	}
	fmt.Println("moved directory", oldStaticDir, " to ", newStaticDir)
	return nil
}
