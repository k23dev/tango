package templates

import (
	"log"
	"strings"
	"tango_cli/parser"
)

type TemplatesFunc struct {
	Name string
	Fn   func()
}

// README
// Para reemplazar los archivos tienen una connotación especial
// $[TIPO: Singular | Plural][Capitalized | Lowercase]$
// EJ: $PC$ (Plural Capitalized)

type Templates struct {
	Parser       *parser.Parser
	Replacements *strings.Replacer
	// Files        map[string][]TemplatesFunc
	Files map[string][]string
}

func New(p *parser.Parser) *Templates {
	t := &Templates{
		Parser: p,
		Files:  make(map[string][]string),
	}
	t.LoadTemplatesFiles()
	return t
}

func (t *Templates) LoadTemplatesFiles() {
	// models

	// TODO: ver la forma de poder hacerlo todo desde una forma diferente

	// tf := &TemplatesFunc{}

	// tf.Name = "default"
	// tf.Fn = models_default(t)

	t.Files["models"] = append(t.Files["models"], "default")
	t.Files["models"] = append(t.Files["models"], "api")
	// tf.Name = "templ"
	// tf.Fn = models_default
	// t.Files["models"] = append(t.Files["models"], *tf)
	//  features
	t.Files["features"] = append(t.Files["features"], "default")
	t.Files["features"] = append(t.Files["features"], "api")
	t.Files["features"] = append(t.Files["features"], "selector")
	t.Files["features"] = append(t.Files["features"], "api")
	// views
	t.Files["views"] = append(t.Files["views"], "default")
	t.Files["views"] = append(t.Files["views"], "selector")
	// routes
	t.Files["routes"] = append(t.Files["routes"], "default")
	t.Files["routes"] = append(t.Files["routes"], "api")

}

func (t *Templates) setReplacements() {

	pc := t.Parser.ConvertToTitle(t.Parser.NamePlural)
	pl := t.Parser.NamePlural
	sc := t.Parser.ConvertToTitle(t.Parser.NameSingular)
	sl := t.Parser.NameSingular
	fl := t.Parser.FirstChar

	t.Replacements = strings.NewReplacer("$PC$", pc, "$PL$", pl, "$SC$", sc, "$SL$", sl, "$FL$", fl)

}

func (t *Templates) SelectAndParse(templatetype, templatefile string) {

	var template string
	templates := t.Files[templatetype]

	for _, file := range templates {
		if file == templatefile {

		}
	}

	log.Printf("%s \n", template)
}
