package locator

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type OfficeRepo struct {
}

func (OfficeRepo) LoadOffices(path string) ([]Office, error) {
	offices := []Office{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return []Office{}, fmt.Errorf("Unable to load offices from: %s", path)
	}

	_ = yaml.Unmarshal(data, &offices)
	return offices, nil
}
