package config

import (
	"fmt"
)

var knownSidecars = map[string]struct{}{"google-cloud-sql-proxy": struct{}{}}

// Sidecar represents a sidecar that can be applied to a grift, migration, or web app
type Sidecar string

// Validate returns nil if s is a valid error
func (s *Sidecar) Validate(m *Meta) error {
	_, ok := knownSidecars[string(*s)]
	if !ok {
		return fmt.Errorf("The sidecar %s isn't supported yet! Maybe file an issue?", s)
	}
	return nil
}
