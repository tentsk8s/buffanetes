package helm

import (
	"fmt"
)

func ChartYAML(org, app string) string {
	// TODO: home, version, description, maintainers
	return fmt.Sprintf(`name: %s-%s
home: buffanetes
version: 0.0.1
description: App created with Buffanetes
maintainers:
- Aaron Schlesinger
`, org, app)
}
