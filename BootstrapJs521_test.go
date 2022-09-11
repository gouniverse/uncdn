package uncdn

import (
	"strings"
	"testing"
)

func TestBootstrapJs521(t *testing.T) {
	output := BootstrapJs521()
	expected := "data-bs-no-jquery"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
