package uncdn

import (
	"strings"
	"testing"
)

func TestJqueryUiJs1332(t *testing.T) {
	output := JqueryUiJs1132()
	expected := "jQuery UI - v1.13.2"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqueryCssJs1332(t *testing.T) {
	output := JqueryUiCss1132()
	expected := "jQuery UI - v1.13.2"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
