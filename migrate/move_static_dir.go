package migrate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func MoveStaticDir(slug string) (err error) {
	year := YearStrip(slug)
	city := CityStrip(slug)
	oldStaticDir := filepath.Join(GetOldEventStaticPath(city, year), "events", slug)
	oldStaticDir += "/"
	newStaticDir := filepath.Join(GetNewEventStaticPath(city, year), "events", year, city)
	newStaticDir += "/"
	if _, err := os.Stat(newStaticDir); err == nil {
		fmt.Println("New directory already exists")
		return nil
	}
	err = os.MkdirAll(filepath.Join(GetNewEventStaticPath(city, year), "events", year), 0755)
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
