package render

import (
	"fmt"
	"os"
	pathutils "path"
	"regexp"
	"strings"

	"github.com/bejelith/docker_entrypoint/pkg/template"
)

var templateExtension = ".template"

func ExecTemplates(paths ...string) error {
	for _, path := range paths {
		fmt.Println(path)
		if !strings.HasSuffix(path, templateExtension) {
			return fmt.Errorf("%s is not a .template file", path)
		}
		destPath := removeExtension(path)
		envVars := getEnvironVars(replaceWithUnderscore(destPath))
		t := template.New(path, destPath, envVars)
		if err := t.Write(); err != nil {
			return fmt.Errorf("error rendering template %s: %s", path, err)
		}
		fmt.Printf("Template %s rendered in %s\n", path, t.DstPath())
	}
	return nil
}

func replaceWithUnderscore(file string) string {
	replacer := strings.NewReplacer(" ", "_", "?", "_", ".", "_")
	return replacer.Replace(pathutils.Base(file))
}

func removeExtension(path string) string {
	p := path[:len(path)-len(templateExtension)]
	return p
}

func getEnvironVars(filterString string) map[string]string {
	regex, _ := regexp.Compile("^" + filterString + "_")
	vars := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitAfterN(e, "=", 2)
		if regex.MatchString(pair[0]) {
			vars[pair[0][len(filterString)+1:len(pair[0])-1]] = pair[1]
		}
	}
	return vars
}
