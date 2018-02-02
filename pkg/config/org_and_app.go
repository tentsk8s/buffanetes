package config

import (
	"errors"
	"fmt"
	"strings"
)

// OrgAndApp represents an organization and an app
type OrgAndApp struct {
	Org string
	App string
}

func (o *OrgAndApp) String() string {
	return fmt.Sprintf("%s/%s", o.Org, o.App)
}

// WithSuffix returns o.Org-o.App-suf
func (o *OrgAndApp) WithSuffix(suf string) string {
	return fmt.Sprintf("%s-%s", o.WithHyphens(), suf)
}

// WithHyphens returns o.Org-o.App
func (o *OrgAndApp) WithHyphens() string {
	return fmt.Sprintf("%s-%s", o.Org, o.App)
}

// ParseOrgAndApp parses str and translates it into an OrgAndApp. Returns
// nil and an error if str is malformed
func ParseOrgAndApp(str string) (*OrgAndApp, error) {
	spl := strings.Split(str, "/")
	if len(spl) != 2 {
		return nil, errors.New("org and app needs to be written as <org>/<app>")
	}
	return &OrgAndApp{Org: spl[0], App: spl[1]}, nil
}
