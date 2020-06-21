package paths

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// EventContentPath returns the path for content for an event based upon city and year
func EventContentPath(city, year string) (eventContentPath string) {
	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)}
	if err := os.MkdirAll(filepath.Join(GetWebdir(), "content", "events", strings.Join(s, "")), 0777); err != nil {
		log.Fatal(err)
	}
	return filepath.Join(GetWebdir(), "content", "events", strings.Join(s, ""))
}

// EventNewContentPath returns the path for content for an event based upon city and year
func EventNewContentPath(city, year string) (eventContentPath string) {
	// if err := os.MkdirAll(filepath.Join(GetWebdir(), "content", "events", year, city)), 0777); err != nil {
	// 	log.Fatal(err)
	// }
	return filepath.Join(filepath.Join(GetWebdir(), "content", "new-events", year, city))
}
