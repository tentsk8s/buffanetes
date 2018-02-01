package config

import (
	"github.com/BurntSushi/toml"
)

// Meta represets the meta.toml file
type Meta struct {
	Version int               `toml:"version"`
	Env     map[string]string `toml:"env"`
}

// ParseMeta parses metadata from a byte slice
func ParseMeta(b []byte) (*Meta, error) {
	ret := new(Meta)
	if err := toml.Unmarshal(b, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
