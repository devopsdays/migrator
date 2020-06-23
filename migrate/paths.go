package migrate

import (
	"path/filepath"
	"strings"
)

func GetOldWebDir() string {
	return "/Users/mattstratton/src/github.com/devopsdays/devopsdays-web"
}

func GetNewWebDir() string {
	return "/Users/mattstratton/src/migrate/devopsdays-web"
}

// GetOldEventStaticPath returns the full path of the old static directory for an event
func GetOldEventStaticPath(city, year string) (eventStaticPath string) {
	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)}
	eventStaticPath = filepath.Join(GetOldWebDir(), "static", "events", strings.Join(s, ""))
	return eventStaticPath
}

// GetNewEventStaticPath returns the full path of the new static directory for an event
func GetNewEventStaticPath(city, year string) (eventStaticPath string) {
	eventStaticPath = filepath.Join(GetNewWebDir(), "static", "events", year, city)
	return eventStaticPath
}

// GetOldEventContentPath returns the old path for content for an event based upon city and year
func GetOldEventContentPath(city, year string) (oldEventContentPath string) {
	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)}
	return filepath.Join(GetOldWebDir(), "content", "events", strings.Join(s, ""))
}

// GetNewEventContentPath returns the new path for content for an event based upon city and year
func GetNewEventContentPath(city, year string) (newEventContentPath string) {
	return filepath.Join(filepath.Join(GetNewWebDir(), "content", "new-events", year, city))
}

// GetOldSponsorsPath returns the old path for sponsor data files
func GetOldSponsorsPath() string {
	return filepath.Join(GetOldWebDir(), "data", "sponsors")

}

// GetNewSponsorsPath returns the new path for sponsor content files
func GetNewSponsorsPath() string {
	return filepath.Join(GetNewWebDir(), "content", "sponsors")

}

// GetEventDataFilePath returns the full path the the data directory for events
func GetEventDataFilePath(city, year string) (eventDataFilePath string) {
	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10), ".yml"}
	eventDataFilePath = filepath.Join(GetOldWebDir(), "data", "events", strings.Join(s, ""))
	// eventDataPath = strings.Join(s, "")
	// eventDataPath = webdir
	return eventDataFilePath
}
