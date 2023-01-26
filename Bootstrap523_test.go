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
		"cosmo":     BootstrapCosmoCss523,
		"cyborg":    BootstrapCyborgCss523,
		"darkly":    BootstrapDarklyCss523,
		"flatly":    BootstrapFlatlyCss523,
		"journal":   BootstrapJournalCss523,
		"litera":    BootstrapLiteraCss523,
		"lumen":     BootstrapLumenCss523,
		"lux":       BootstrapLuxCss523,
		"materia":   BootstrapMateriaCss523,
		"minty":     BootstrapMintyCss523,
		"morph":     BootstrapMorphCss523,
		"pulse":     BootstrapPulseCss523,
		"quartz":    BootstrapQuartzCss523,
		"sandstone": BootstrapSandstoneCss523,
		"simplex":   BootstrapSimplexCss523,
		"sketchy":   BootstrapSketchyCss523,
		"slate":     BootstrapSlateCss523,
		"solar":     BootstrapSolarCss523,
		"spacelab":  BootstrapSpacelabCss523,
		"superhero": BootstrapSuperheroCss523,
		"united":    BootstrapUnitedCss523,
		"vapor":     BootstrapVaporCss523,
		"yeti":      BootstrapYetiCss523,
		"zephyr":    BootstrapZephyrCss523,
	}
	for key, _ := range table {
		output := table[key]()
		expected := ".container,.container-lg,.container-md,.container-sm,.container-xl"
		if !strings.Contains(output, expected) {
			t.Error("Does not contain '" + expected + "', Output:" + output)
		}
	}
}
