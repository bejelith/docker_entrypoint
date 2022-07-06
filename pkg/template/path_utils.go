package template

import (
	"os"
	pathutils "path"
	"regexp"
	"strings"
)

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
