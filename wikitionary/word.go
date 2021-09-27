package wikitionary

import (
	"bufio"
	"io"
	"strings"

	"github.com/dvoeglazyi/gospelling/templates"
)

type Word struct {
	Value         string // слово
	PartType      string // часть речи
	Template      *templates.Template
	Bases         map[string]string
	Transcription struct {
		Variants     []string // варианты
		AccentedChar int      // номер буквы под ударением
	}
	Morphological struct {
		Root   string // корень
		Suffix string // суффикс
	}
}

func (p *Parser) getWord(s string) (*Word, error) {
	r := bufio.NewReader(strings.NewReader(s))
	word := Word{}
	var reminder, template string
	var err error
	for {
		if reminder, template, err = readTag(r, reminder); err != nil {
			if err == io.EOF {
				return nil, nil
			}
			return nil, err
		}
		template = trimTag(template)

		if word.PartType = getPartType(template); word.PartType == "" {
			continue
		}
		hasTemplate := false
		for _, t := range p.Templates {
			if strings.HasPrefix(template, t.Title) {
				word.Template = &t
				hasTemplate = true
				break
			}
		}
		if !hasTemplate {
			continue
		}
		word.Bases = GetWordBasesFromTag(template)
		// TODO при получении других тэгов прийдётся сделать выход из цикла иначе
		return &word, nil
	}
}

func getPartType(s string) string {
	if !strings.Contains(s, "ru") {
		return ""
	}
	for _, speechPart := range templates.SpeechParts {
		if strings.HasPrefix(s, speechPart) {
			return speechPart
		}
	}
	return ""
}
