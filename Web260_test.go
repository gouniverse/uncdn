package uncdn

import (
	"strings"
	"testing"
)

func TestWeb260(t *testing.T) {
	output := Web260()
	expected := "Config"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
