package config

import (
	"fmt"
	"io/ioutil"

	"github.com/DatTran1597/golang-starter/model"
	yaml "gopkg.in/yaml.v2"
)

func Load(path string) (*model.Configuration, error) {
	cfg := &model.Configuration{}
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file , %s", err)
	}

	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return cfg, nil
}
