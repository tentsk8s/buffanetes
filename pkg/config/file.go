package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

func ParseFile(name string, out interface{}) error {
	full := filepath.Join(".buffanetes", name)
	b, err := ioutil.ReadFile(full)
	if err != nil {
		return err
	}
	if err := toml.Unmarshal(b, out); err != nil {
		return err
	}
	return nil
}
