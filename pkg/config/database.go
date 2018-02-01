package config

import (
	"github.com/BurntSushi/toml"
)

// Databases is the top-level databases object that is seen in databases.toml
type Databases map[string]Database

// Database is details of an individual database
type Database struct {
	Host               string `toml:"host"`
	Port               int    `toml:"port"`
	Username           string `toml:"username"`
	PasswordSecretName string `toml:"password-secret"`
}

func ParseDatabases(b []byte) (*Databases, error) {
	ret := new(Databases)
	if err := toml.Unmarshal(b, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
