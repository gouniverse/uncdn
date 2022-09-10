package uncdn

import (
	"strings"
	"testing"
)

func TestWebJs260(t *testing.T) {
	output := WebJs260()
	expected := "Config"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
