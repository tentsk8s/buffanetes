package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Web is the top level config directive that tells Buffanetes how to deploy the web app
type Web struct {
	HostEnv  string    `toml:"host-env"`
	Database string    `toml:"database"`
	Sidecars []Sidecar `toml:"sidecars"`
	Envs     []WebEnv  `toml:"env"`
	Health   Health    `toml:"health"`
}

func ParseWeb(b []byte) (*Web, error) {
	ret := new(Web)
	if err := toml.Unmarshal(b, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// WebEnv is the config directive that tells Buffanetes how to inject an env var into
// the web app
type WebEnv struct {
	From       string  `toml:"from"`
	Name       string  `toml:"name"`
	SecretName *string `toml:"name"`
}

// Validate returns nil if w is a valid config
func (w *WebEnv) Validate(m *Meta) error {
	if w.From == "secret" && w.SecretName == nil {
		return fmt.Errorf("The env var %s doesn't have a secret name", w.Name)
	}
	return nil
}

// Validate returns nil if w is a valid configuration
func (w *Web) Validate(m *Meta) error {
	for _, env := range w.Envs {
		if err := env.Validate(m); err != nil {
			return err
		}
	}
	for _, sidecar := range w.Sidecars {
		if err := sidecar.Validate(m); err != nil {
			return err
		}
	}
	// TODO: database validation
	return nil
}
