package uncdn

import (
	"strings"
	"testing"
)

func TestSweetalert2_11432(t *testing.T) {
	output := Sweetalert2_11432()
	expected := "SweetAlert2"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
