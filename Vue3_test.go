package uncdn

import (
	"strings"
	"testing"
)

func TestVue3(t *testing.T) {
	output := Vue3()
	expected := "Vue"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
