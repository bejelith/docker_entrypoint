package template

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var baseVar string
var testVar string
var testTemplate []byte
var tempDir string

func TestMain(m *testing.M) {
	testVar = "AAA"
	baseVar = "X"
	testTemplate = []byte(fmt.Sprintf("a={{.%s}}", testVar))
	tempDir, _ = ioutil.TempDir("", "gotest.ext")
	m.Run()
	_ = os.Remove(tempDir)
}

func mockFile(content []byte) (string, error) {
	if tempFile, err := os.Create(tempDir + "/file.temp"); err == nil {
		defer func() { _ = tempFile.Close() }()
		_, err = tempFile.Write(content)
		return tempFile.Name(), err
	} else {
		return "", err
	}
}

func TestTemplate(t *testing.T) {
	var err error
	var templatePath string
	var resultPath = tempDir + "/result"
	testVal := "BBB"

	templatePath, err = mockFile(testTemplate)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		testVar: testVal,
	}
	template := New(templatePath, templatePath, vars)
	if err := template.WriteToPath(resultPath); err != nil {
		t.Fatal(err)
	}
	if file, err := ioutil.ReadFile(resultPath); err == nil {
		if string(file) != "a="+testVal {
			t.Fatalf("File is %s, expected %s", string(file), "a="+testVal)
		}
	} else {
		t.Fatal(err)
	}
}

var trimCases = map[string]string{
	".extension":     "",
	"file.extension": "file",
	"file-extension": "file-extension",
	"file.":          "file",
}

func TestTrimExtension(t *testing.T) {
	for k, v := range trimCases {
		if trimExtension(k) != v {
			t.Fatalf("Received %s, expected %s", k, v)
		}
	}
}
