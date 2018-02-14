package web

import (
	"github.com/tentsk8s/buffanetes/pkg/config"
	"k8s.io/api/core/v1"
)

// Deployment returns the web deployment for orgAndApp
func Deployment(orgAndApp *config.OrgAndApp) *v1.Deployment {
	return &v1.Deployment{}
}
