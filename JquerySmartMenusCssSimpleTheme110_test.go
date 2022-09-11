package uncdn

import (
	"strings"
	"testing"
)

func TestJquerySmartMenusCssSimpleTheme110(t *testing.T) {
	output := JquerySmartMenusCssSimpleTheme110()
	expected := ".sm"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
