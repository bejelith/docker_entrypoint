package render

import (
	"fmt"

	"github.com/bejelith/docker_entrypoint/pkg/template"
)

func ExecTemplates(templates ...*template.Template) error {
	for _, t := range templates {
		if err := t.Write(); err != nil {
			return fmt.Errorf("error rendering template %s: %s", t.DstPath(), err)
		}
		fmt.Printf("Template rendered in %s\n", t.DstPath())
	}
	return nil
}