package migrate

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func MoveStaticDir(slug string) (err error) {
	year := YearStrip(slug)
	city := CityStrip(slug)
	oldStaticDir := GetOldEventStaticPath(city, year)
	oldStaticDir += "/"
	newStaticDir := GetNewEventStaticPath(city, year)
	newStaticDir += "/"
	// if _, err := os.Stat(newStaticDir); err == nil {
	// 	fmt.Println("New directory already exists")
	// 	return nil
	// }
	err = os.MkdirAll(filepath.Join(GetNewWebDir(), "static", "events", year), 0755)
	if err != nil {
		return errors.Wrap(err, "make year static directory failed")
	}
	fmt.Println("Created new static year directory at ", filepath.Join(GetNewWebDir(), "static", "events", year))
	// err = os.Rename(oldStaticDir, newStaticDir)
	// fmt.Println("Old static directory is ", oldStaticDir)
	// fmt.Println("New static directory is ", newStaticDir)
	// if err != nil {
	// 	return errors.Wrap(err, "rename directory failed")
	// }
	// fmt.Println("moved directory", oldStaticDir, " to ", newStaticDir)

	err = copy_folder(oldStaticDir, newStaticDir)
	if err != nil {
		return errors.Wrap(err, "copy static folder failed")
	}
	// TODO: delete old static/events/yyyy-city directory in GetNewWebDir()
	return nil
}

func copy_folder(source string, dest string) (err error) {

	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			err = copy_folder(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = copy_file(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func copy_file(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}
