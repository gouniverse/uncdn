package uncdn

import (
	"strings"
	"testing"
)

func TestJquery360(t *testing.T) {
	output := Jquery360()
	expected := "jQuery v3.6.0"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
