package web

import (
	"fmt"
)

func ServiceYAML(org, app string) string {
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
`, baseName(org, app), baseName(org, app))
}
