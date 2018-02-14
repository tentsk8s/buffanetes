package helm

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// // BuildChart outputs a helm chart from chart to dirName. dirName must be an
// // absolute path and already exist
// func BuildChart(dir string, chart *Chart) error {
// 	// build the template
// 	tplDir := filepath.Join(dir, "templates")
// 	if err := os.MkdirAll(tplDir, os.ModePerm); err != nil {
// 		return err
// 	}

// 	// write Chart.yaml, README.md and values.yaml
// 	if err := ioutil.WriteFile(
// 		filepath.Join(dir, "Chart.yaml"),
// 		[]byte(chart.ChartYAML),
// 		os.ModePerm,
// 	); err != nil {
// 		return err
// 	}

// 	if err := ioutil.WriteFile(
// 		filepath.Join(dir, "README.md"),
// 		[]byte(chart.Readme),
// 		os.ModePerm,
// 	); err != nil {
// 		return err
// 	}

// 	valuesYAML, err := yaml.Marshal(chart.Values)
// 	if err != nil {
// 		return err
// 	}

// 	if err := ioutil.WriteFile(
// 		filepath.Join(dir, "values.yaml"),
// 		valuesYAML,
// 		os.ModePerm,
// 	); err != nil {
// 		return err
// 	}

// 	// now write all the templates
// 	for _, tpl := range chart.Templates {
// 		if err := writeTemplate(tplDir, tpl); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func writeTemplate(baseDir string, tpl *Template) error {
// 	return ioutil.WriteFile(
// 		filepath.Join(baseDir, tpl.Filename),
// 		[]byte(tpl.Contents),
// 		os.ModePerm,
// 	)
// }
