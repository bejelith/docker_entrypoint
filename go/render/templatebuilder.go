package render

import (
	"fmt"
	"os"
	pathutils "path"
	"regexp"
	"strings"

	"github.com/bejelith/docker_entrypoint/template"
)

var templateExtension = ".template"

func ExecTemplates(paths ...string) error {
	for _, path := range paths {
		if !strings.HasSuffix(path, templateExtension) {
			return fmt.Errorf("%s is not a .template file", path)
		}
		destPath := removeExtension(path)

		filterString := strings.ToUpper(replaceWithUnderscore(destPath))
		envVars, err := getEnvironVars(filterString)
		if err != nil {
			return fmt.Errorf("error getting environment variables for %s: %s", filterString, err)
		}

		t := template.New(path, destPath, envVars)
		if err := t.Write(); err != nil {
			return fmt.Errorf("error rendering template %s: %s", path, err)
		}
		fmt.Printf("Template %s rendered\n", path)
	}
	return nil
}

func replaceWithUnderscore(file string) string {
	replacer := strings.NewReplacer(" ", "_", "?", "_", ".", "_")
	return replacer.Replace(pathutils.Base(file))
}

func removeExtension(path string) string {
	return path[:len(path)-len(templateExtension)]
}

func getEnvironVars(filterString string) (map[string]string, error) {
	prefix := fmt.Sprintf("%s_", filterString)
	filter := fmt.Sprintf("^%s", prefix)
	regex, err := regexp.Compile(filter)

	if err != nil {
		return nil, err
	}

	vars := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if regex.MatchString(pair[0]) {
			vars[pair[0][len(prefix):]] = pair[1]
		}
	}
	return vars, nil
}
