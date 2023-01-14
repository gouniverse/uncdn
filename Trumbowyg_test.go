package uncdn

import (
	"strings"
	"testing"
)

func TestTrumbowyg2260Css(t *testing.T) {
	output := Trumbowyg2260Css()
	expected := "Trumbowyg"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestTrumbowyg2260Js(t *testing.T) {
	output := Trumbowyg2260Js()
	expected := "Trumbowyg"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
