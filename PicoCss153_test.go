package uncdn

import (
	"strings"
	"testing"
)

func TestPicoCss153(t *testing.T) {
	output := PicoCss153()
	expected := "Pico.css v1.5.3"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
