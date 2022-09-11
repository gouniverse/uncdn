package uncdn

import (
	"strings"
	"testing"
)

func TestJquerySmartMenusCssBootstrap4AddOn110(t *testing.T) {
	output := JquerySmartMenusCssBootstrap4AddOn110()
	expected := ".navbar-nav.sm-collapsible"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
