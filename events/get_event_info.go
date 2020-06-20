package events

import (
	"io/ioutil"

	"github.com/devopsdays/migrator/model"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

func GetEventInfo(file string) (event model.EventData, err error) {

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return event, errors.Wrapf(err, "load event YAML file failed for %s", file)
	}

	err = yaml.Unmarshal(yamlFile, &event)
	if err != nil {
		panic(err)
	}

	return
}
