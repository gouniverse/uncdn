package uncdn

import (
	"strings"
	"testing"
)

func TestJquerySmartMenusCss110(t *testing.T) {
	output := JquerySmartMenusCss110()
	expected := ".sm"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
