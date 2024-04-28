package templates

import (
	"strings"
	"tango_cli/parser"
)

// README
// Para reemplazar los archivos tienen una connotación especial
// $[TIPO: Singular | Plural][Capitalized | Lowercase]$
// EJ: $PC$ (Plural Capitalized)

type Templates struct {
	Parser       *parser.Parser
	Replacements *strings.Replacer
}

func New(p *parser.Parser) *Templates {
	return &Templates{
		Parser: p,
	}
}

func (t *Templates) setReplacements() {

	pc := t.Parser.ConvertToTitle(t.Parser.NamePlural)
	pl := t.Parser.NamePlural
	sc := t.Parser.ConvertToTitle(t.Parser.NameSingular)
	sl := t.Parser.NameSingular
	fl := t.Parser.FirstChar

	t.Replacements = strings.NewReplacer("$PC$", pc, "$PL$", pl, "$SC$", sc, "$SL$", sl, "$FL$", fl)

}
