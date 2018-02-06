package web

import (
	"fmt"

	"github.com/tentsk8s/buffanetes/pkg/config"
)

func ServiceYAML(orgAndApp *config.OrgAndApp) string {
	return fmt.Sprintf(`apiVersion: v1
kind: Service
metadata:
	name: %s
spec:
	type: ClusterIP
	ports:
	- name: http
	  port: 8080
	  targetPort: 8080
	selector:
		app: %s
`, orgAndApp.WithHyphens(""), orgAndApp.WithHyphens())
}
