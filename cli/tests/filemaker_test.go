package tests

import (
	"strings"
	"tango_cli/filemaker"
	"testing"
)

func TestFileMakerSetRoot(t *testing.T) {

	fm := filemaker.New("app")

	fm.SetRootPath("")

	gotBuffer := strings.Split(fm.RootPath, "/")
	path := "/tmp/randomfolder/other"
	wantBuffer := strings.Split(path, "/")
	got := len(gotBuffer)
	want := len(wantBuffer)

	if got != want {
		t.Errorf("Tengo %d | Quiero %d", got, want)
	} else {
		t.Logf("Tengo %d | Quiero %d", got, want)
	}

}

func TestFileMakerFileNotExists(t *testing.T) {

	fm := filemaker.New("app")

	fm.SetRootPath("")

	filepath := "app/models/category.go"

	got := fm.CheckIfFileExists(filepath)

	want := false

	if got != want {

		t.Errorf("Tengo %v | Quiero %v", got, want)
	} else {
		t.Logf("Tengo %v | Quiero %v", got, want)
	}

}
