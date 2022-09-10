package uncdn

package uncdn

import (
	"strings"
	"testing"
)

func TestPhpStrToTimeJs20220910_test(t *testing.T) {
	output := PhpStrToTimeJs20220910()
	expected := "strtotime"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
