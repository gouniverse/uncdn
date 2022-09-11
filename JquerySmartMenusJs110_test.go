package uncdn

import (
	"strings"
	"testing"
)

func TestJquerySmartMenusJs110(t *testing.T) {
	output := JquerySmartMenusJs110()
	expected := ".container,.container-lg,.container-md,.container-sm,.container-xl"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
