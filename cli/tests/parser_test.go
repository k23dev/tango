package tests

import (
	"tango_cli/parser"
	"testing"
)

func TestParserPlural(t *testing.T) {

	p := parser.New()
	p.Read("ArChivo")

	got := p.NamePlural
	want := "archivos"

	if got != want {
		t.Errorf("Tengo %v | Quiero %v", got, want)
		t.Fail()
	} else {
		t.Logf("Tengo %v | Quiero %v", got, want)
	}

}

func TestParserPlural2(t *testing.T) {

	p := parser.New()
	p.Read("Directory")

	got := p.NamePlural
	want := "directories"

	if got != want {
		t.Errorf("Tengo %v | Quiero %v", got, want)
		t.Fail()
	} else {
		t.Logf("Tengo %v | Quiero %v", got, want)
	}

}
func TestParserSingular(t *testing.T) {

	p := parser.New()
	p.Read("Diarios")

	got := p.NameSingular
	want := "diario"

	if got != want {
		t.Errorf("Tengo %v | Quiero %v", got, want)
		t.Fail()
	} else {
		t.Logf("Tengo %v | Quiero %v", got, want)
	}
}

func TestParserFirstChar(t *testing.T) {

	p := parser.New()
	p.Read("Chuleta")

	got := p.FirstChar
	want := "c"

	if got != want {
		t.Errorf("Tengo %v | Quiero %v", got, want)
		t.Fail()
	} else {
		t.Logf("Tengo %v | Quiero %v", got, want)
	}
}

func TestParserTitle(t *testing.T) {

	p := parser.New()
	p.Read("perritos")

	got := p.ConvertToTitle(p.NameSingular)
	want := "Perrito"

	if got != want {
		t.Errorf("Tengo %v | Quiero %v", got, want)
		t.Fail()
	} else {
		t.Logf("Tengo %v | Quiero %v", got, want)
	}
}
