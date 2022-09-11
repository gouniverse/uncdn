package uncdn

import (
	"strings"
	"testing"
)

func TestJquerySmartMenusCssBlueTheme110(t *testing.T) {
	output := JquerySmartMenusCssBlueTheme110()
	expected := ".sm"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
