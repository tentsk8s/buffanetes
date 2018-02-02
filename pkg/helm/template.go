package helm

import ()

// Template represents a single helm template
type Template struct {
	Filename string
	Contents string
}

func (t *Template) String() string {
	return t.Filename
}
