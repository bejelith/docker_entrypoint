package template

import (
	"os"
	"path/filepath"
	gotemplate "text/template"
)

type Template struct {
	vars       map[string]string
	sourcePath string
	destPath   string
}

func New(templatePath string, outputPath string, vars map[string]string) Template {
	t := Template{
		vars:       vars,
		sourcePath: templatePath,
		destPath:   outputPath,
	}
	return t
}

func (t *Template) Write() error {
	return t.WriteToPath(t.destPath)
}

func (t *Template) WriteToPath(path string) error {
	var err error
	if file, osErr := os.Create(path); osErr == nil {
		defer file.Close()
		if template, parseError := gotemplate.ParseFiles(t.sourcePath); parseError == nil {
			template.Option("missingkey=error")
			return template.Execute(file, t.vars)
		} else {
			err = parseError
		}
	} else {
		err = osErr
	}
	return err
}

func (t *Template) DstPath() string {
	return t.destPath
}

func trimExtension(path string) string {
	var extension = filepath.Ext(path)
	return path[0 : len(path)-len(extension)]
}
