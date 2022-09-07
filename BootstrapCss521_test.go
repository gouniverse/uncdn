package uncdn

import (
	"strings"
	"testing"
)

func TestBootstrapCss521(t *testing.T) {
	output := BootstrapCss521()
	expected := "Bootstrap  v5.2.1"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
