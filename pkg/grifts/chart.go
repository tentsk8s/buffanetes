package grifts

import (
	"k8s.io/helm/pkg/proto/hapi/chart"
)

func BuildChart(imageName, imageTag, griftName string) *chart.Chart {
	return &chart.Chart{}
}
