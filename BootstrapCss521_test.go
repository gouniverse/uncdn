package uncdn

import (
	"strings"
	"testing"
)

func TestBootstrapCss521(t *testing.T) {
	output := BootstrapCss521()
	expected := ".container,.container-lg,.container-md,.container-sm,.container-xl"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
