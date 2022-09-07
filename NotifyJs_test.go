package uncdn

import (
	"strings"
	"testing"
)

func TestNotifyJs(t *testing.T) {
	output := NotifyJs()
	expected := "Notify"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
