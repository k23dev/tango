package parser

import (
	"fmt"
	"strings"
)

type Parser struct {
	nameInput    string
	NamePlural   string `json:"name_plural"`
	NameSingular string `json:"name_singular"`
	FirstChar    string `json:"first_char"`
	LastCharPos  int
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Read(name string) {

	p.nameInput = name

	// Convert the input
	p.convertToPlural()
	p.convertToSingular()
	p.getFirstChar()

}

func (p *Parser) convertToPlural() {

	// todo
	// si termina en vocal agrega S
	// si termina en Y agrega "ies"
	p.LastCharPos = len(p.nameInput) - 1
	lastChar := string(p.nameInput[p.LastCharPos])
	pluralEnd := "s"

	buf := p.nameInput
	if lastChar == "y" {
		pluralEnd = "ies"
		buf = p.nameInput[0:p.LastCharPos]
	} else if lastChar == "s" {
		pluralEnd = ""
	}

	p.NamePlural = fmt.Sprintf("%s%s", buf, pluralEnd)
	p.NamePlural = strings.ToLower(p.NamePlural)
}

func (p *Parser) convertToSingular() {
	if string(p.nameInput[p.LastCharPos]) == "s" {
		p.NameSingular = p.nameInput[0:p.LastCharPos]
	} else {
		p.NameSingular = p.nameInput
	}
	p.NameSingular = strings.ToLower(p.NameSingular)
}

func (p *Parser) ConvertToTitle(s string) string {
	return strings.Title(s)
}

func (p *Parser) getFirstChar() {
	p.FirstChar = string(p.nameInput[0:1])
	p.FirstChar = strings.ToLower(p.FirstChar)
}
