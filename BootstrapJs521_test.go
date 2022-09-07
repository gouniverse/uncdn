package uncdn

import (
	"strings"
	"testing"
)

func TestBootstrapJs521(t *testing.T) {
	output := BootstrapJs521()
	expected := "Bootstrap v5.2.1"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
