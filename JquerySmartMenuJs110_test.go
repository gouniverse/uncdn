package uncdn

import (
	"strings"
	"testing"
)

func TestJquerySmartMenuJs110(t *testing.T) {
	output := JquerySmartMenuJs110()
	expected := ".container,.container-lg,.container-md,.container-sm,.container-xl"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
