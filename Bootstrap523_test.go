package uncdn

import (
	"strings"
	"testing"
)

func TestBootstrapCss523(t *testing.T) {
	output := BootstrapCss523()
	expected := ".container,.container-lg,.container-md,.container-sm,.container-xl"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapJs523(t *testing.T) {
	output := BootstrapJs523()
	expected := "data-bs-no-jquery"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapCeruleanCss523(t *testing.T) {
	table := map[string]func() string{
		"bootstrap": BootstrapCeruleanCss523,
		"cerulean":  BootstrapCeruleanCss523,
		"yeti":      BootstrapYetiCss523,
	}
	for key, _ := range table {
		output := table[key]()
		expected := ".container,.container-lg,.container-md,.container-sm,.container-xl"
		if !strings.Contains(output, expected) {
			t.Error("Does not contain '" + expected + "', Output:" + output)
		}
	}
}
