package uncdn

import (
	"strings"
	"testing"
)

func TestPhpDateJs20220910_test(t *testing.T) {
	output := PhpDateJs20220910()
	expected := "date"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
