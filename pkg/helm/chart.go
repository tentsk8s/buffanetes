package helm

import (
	"github.com/tentsk8s/buffanetes/pkg/config"
	"k8s.io/helm/pkg/proto/hapi/chart"
)

func chartMetadata(orgAndApp config.OrgAndApp) *chart.Metadata {
	return &chart.Metadata{
		Name:        orgAndApp.String(),
		Version:     "v0.0.1",
		Description: orgAndApp.String(),
		Maintainers: []*chart.Maintainer{
			{
				Name:  "buffanetes",
				Email: "info@buffanetes.io",
				Url:   "https://github.com/tentsk8s/buffanetes",
			},
		},
	}
}

// NewChart creates a new *Chart from the given parameters
func NewChart(orgAndApp config.OrgAndApp) *chart.Chart {
	return &chart.Chart{
		Metadata: chartMetadata(orgAndApp),
	}
}
