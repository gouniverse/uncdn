package uncdn

import (
	"strings"
	"testing"
)

func TestVueJs3(t *testing.T) {
	output := VueJs3()
	expected := "Vue"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
