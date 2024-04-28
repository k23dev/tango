package filemaker

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"tango_cli/parser"
	"tango_cli/templates"
)

type FileMaker struct {
	Filename         string   `json:"filename"`
	Filetypes        []string `json:"filetypes"`
	RootPath         string   `json:"root_path"`
	Mode             string   `json:"mode"`
	ForcedMode       bool     `json:"forced_mode"`
	AppDir           string   `json:"app_dir"`
	TemplateSelected string   `json:"template_selected"`
	FilePerms        fs.FileMode
	Parser           *parser.Parser
	Templates        *templates.Templates
}

func New(packageName string) *FileMaker {
	p := parser.New()
	p.Read(packageName)
	t := templates.New(p)
	t.Parser = p
	return &FileMaker{
		ForcedMode: false,
		FilePerms:  0666,
		Parser:     p,
		Templates:  t,
	}
}

func (fm *FileMaker) setFileTypes() {

	ftypes := []string{
		"templ",
		"go",
	}

	fm.Filetypes = ftypes

}
func (fm *FileMaker) setFilename(filename string) {
	fm.Filename = strings.ToLower(filename)
}

func (fm *FileMaker) SetRootPath(path string) {
	if path != "" {
		fm.RootPath = path + "/"
		return
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fm.RootPath = filepath.Dir(ex) + "/"
}

func (fm *FileMaker) SetMode(mode string) {
	fm.Mode = strings.ToLower(mode)
}

func (fm *FileMaker) SetForcedMode(isForced bool) {
	fm.ForcedMode = isForced
}

func (fm *FileMaker) SetAppDir(dir string) {
	fm.AppDir = dir
}

func (fm *FileMaker) CheckIfFileExists(filepath string) bool {

	if _, err := os.Stat(fm.RootPath + filepath); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}

}

func (fm *FileMaker) MakeIt() {
	// TODO Comprobar si el archivo existe si existe pide lo que haga en modo forzoso. De ser así se eliminan los archivos y se crean estos nuevos.
	switch fm.Mode {
	case "basic":
		fm.buildModeBasic()
	case "full":
		fm.buildModeFull()
	case "fullwithselector":
		fm.buildModeFullWithSelector()
	default:
		fmt.Println("El modo seleccionado no es correcto. Puede elegir entre BASIC / FULL / FULL+")
	}
}

func (fm *FileMaker) GetFilePath(directory, extension string, isPlural bool) string {

	filename := fm.GetFilePathFilename("", "", isPlural)
	return fm.RootPath + fm.AppDir + "/" + directory + "/" + filename + "." + extension
}

func (fm *FileMaker) GetFilePathTitleFixs(directory, prefix, posfix, extension string, isPlural bool) string {

	filename := fm.GetFilePathFilename(prefix, posfix, isPlural)

	return fm.RootPath + fm.AppDir + "/" + directory + "/" + filename + "." + extension
}

func (fm *FileMaker) GetFilePathFilename(prefix, posfix string, isPlural bool) string {

	filename := fm.Parser.NameSingular
	if isPlural {
		filename = fm.Parser.NamePlural
	}
	if prefix != "" {
		filename = prefix + filename
	}
	if posfix != "" {
		filename = filename + posfix
	}
	return filename
}

func (fm *FileMaker) selectTemplate(template string) {
	switch template {
	case "model":
		fm.TemplateSelected = fm.Templates.Model()
	case "feature":
		fm.TemplateSelected = fm.Templates.Feature()
	case "route":
		fm.TemplateSelected = fm.Templates.Route()
	case "view":
		fm.TemplateSelected = fm.Templates.View()
	default:
		fm.TemplateSelected = ""
	}
}

func (fm *FileMaker) selectFullTemplate(template string) {

	switch template {
	case "featureselector":
		fm.TemplateSelected = fm.Templates.FeatureSelector()
	case "table":
		fm.TemplateSelected = fm.Templates.ViewsFullTable()
	case "menu":
		fm.TemplateSelected = fm.Templates.ViewsFullMenu()
	case "form":
		fm.TemplateSelected = fm.Templates.ViewsFullForms()
	case "formwithselector":
		fm.TemplateSelected = fm.Templates.ViewsFullFormsWithSelector()
	case "formwithselectorupdate":
		fm.TemplateSelected = fm.Templates.ViewsFullFormsWithSelectorUpdate()
	case "viewselector":
		fm.TemplateSelected = fm.Templates.ViewSelector()
	case "selector":
		fm.TemplateSelected = fm.Templates.ViewsFullSelector()
	case "selectorupdate":
		fm.TemplateSelected = fm.Templates.ViewsFullSelectorUpdate()
	default:
		fm.TemplateSelected = ""
	}

}

func (fm *FileMaker) saveFile(filepath string, content string) bool {
	if err := os.WriteFile(filepath, []byte(content), fm.FilePerms); err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Println("Archivo creado: ", filepath)
		return true
	}
}

func (fm *FileMaker) builder(directory, extension string, isPlural bool) {
	filepath := fm.GetFilePath(directory, extension, isPlural)
	if fm.TemplateSelected != "" {
		// Creo el archivo
		fm.saveFile(filepath, fm.TemplateSelected)
	} else {
		fmt.Println("No hay existe ese template de archivo")
	}
}

func (fm *FileMaker) builderWithFilenameFixes(directory, prefix, posfix, extension string, isPlural bool) {
	filepath := fm.GetFilePathTitleFixs(directory, prefix, posfix, extension, isPlural)
	if fm.TemplateSelected != "" {
		// Creo el archivo
		fm.saveFile(filepath, fm.TemplateSelected)
	} else {
		fmt.Println("No hay existe ese template de archivo")
	}
}

func (fm *FileMaker) buildModeBasic() {
	// todo
	fm.selectTemplate("model")
	fm.builder("models", "go", false)
	fm.selectTemplate("feature")
	fm.builder("features", "go", true)
	fm.selectTemplate("route")
	fm.builder("routes", "go", true)
}

func (fm *FileMaker) buildModeFull() {

	fm.buildModeBasic()
	fm.selectTemplate("view")
	fm.builder("views", "templ", true)
	// menu
	fm.selectFullTemplate("menu")
	fm.builder("views/menus", "templ", true)
	// forms
	fm.selectFullTemplate("form")
	fm.builder("views/forms", "templ", true)
	// tabla
	fm.selectFullTemplate("table")
	fm.builder("views/tables", "templ", true)
}

func (fm *FileMaker) buildModeFullWithSelector() {

	// model
	fm.selectTemplate("model")
	fm.builder("models", "go", false)
	// feature
	fm.selectFullTemplate("featureselector")
	fm.builder("features", "go", true)
	// update
	fm.selectTemplate("route")
	fm.builder("routes", "go", true)
	// view selector
	fm.selectFullTemplate("viewselector")
	fm.builder("views", "templ", true)
	// menu
	fm.selectFullTemplate("menu")
	fm.builder("views/menus", "templ", true)
	// forms
	fm.selectFullTemplate("formwithselector")
	fm.builder("views/forms", "templ", true)
	// form update
	fm.selectFullTemplate("formwithselectorupdate")
	fm.builderWithFilenameFixes("views/forms", "", "_update", "templ", true)
	// tabla
	fm.selectFullTemplate("table")
	fm.builder("views/tables", "templ", true)
	// selector
	fm.selectFullTemplate("selector")
	fm.builderWithFilenameFixes("views/components", "selector_", "", "templ", true)
	// selector update
	fm.selectFullTemplate("selectorupdate")
	fm.builderWithFilenameFixes("views/components", "selector_", "_update", "templ", true)

}
