package wikitionary

import "github.com/dvoeglazyi/gospelling/templates"

type Parser struct {
	Templates []templates.Template
}

func NewParser(templates []templates.Template) *Parser {
	return &Parser{Templates: templates}
}
