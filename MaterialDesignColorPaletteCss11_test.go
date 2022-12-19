package uncdn

import (
	"strings"
	"testing"
)

func TestMaterialDesignColorPaletteCss11(t *testing.T) {
	output := MaterialDesignColorPaletteCss11()
	expected := ".mdc-text-red-A100"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
