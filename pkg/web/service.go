package web

import (
	"fmt"

	"github.com/tentsk8s/buffanetes/pkg/config"
	"k8s.io/api/core/v1"
)

// Service returns the kubernetes service resource for orgAndApp
func Service(orgAndApp *config.OrgAndApp) *v1.Service {
	return &v1.Service{}
	// 	return fmt.Sprintf(`apiVersion: v1
	// kind: Service
	// metadata:
	// 	name: %s
	// spec:
	// 	type: ClusterIP
	// 	ports:
	// 	- name: http
	// 	  port: 8080
	// 	  targetPort: 8080
	// 	selector:
	// 		app: %s
	// `, orgAndApp.WithHyphens(""), orgAndApp.WithHyphens())
}
