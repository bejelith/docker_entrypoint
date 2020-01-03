package render

import (
	"os"
	"testing"
)

func TestGetEnvironVars(t *testing.T) {
	_ = os.Setenv("A_B_C", "1")
	vars, _ := getEnvironVars("A")
	if vars["B_C"] != "1" {
		t.Fatalf("%v", vars)
	}
}
