package helm

import (
	"fmt"

	"github.com/tentsk8s/buffanetes/pkg/config"
)

// Chart represents an entire helm chart
type Chart struct {
	ChartYAML string
	Readme    string
	Values    Values
	Templates []*Template
}

// NewChart creates a new *Chart from the given parameters
func NewChart(chartYAML, readme string, values Values, templates []*Template) *Chart {
	return &Chart{
		ChartYAML: chartYAML,
		Readme:    readme,
		Values:    values,
		Templates: templates,
	}
}

// ChartYAML outputs the chart YAML for the given org and app
func ChartYAML(orgAndApp *config.OrgAndApp) string {
	// TODO: home, version, description, maintainers
	return fmt.Sprintf(`name: %s
home: http://%s.%s
version: 0.0.1
description: %s
maintainers:
- Buffanetes
`, orgAndApp.String(), orgAndApp.Org, orgAndApp.App, orgAndApp.String())
}
