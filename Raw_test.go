package uncdn

import (
	"strings"
	"testing"
)

func TestRaw(t *testing.T) {
	output := Raw("locutus/php/date-20220910.js")
	expected := "Date"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
