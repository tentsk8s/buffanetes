package web

import (
	"github.com/tentsk8s/buffanetes/pkg/config"
	"github.com/tentsk8s/buffanetes/pkg/helm"
)

// BuildHelmChart builds a new helm chart in a temp dir and returns the fully
// qualified path to that temp dir. returns "" and an error if that failed
func BuildHelmChart(webConf *config.Web, orgAndApp *config.OrgAndApp) (string, error) {
	tpls := []*helm.Template{}
}
