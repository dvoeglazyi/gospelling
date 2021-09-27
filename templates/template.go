package templates

import (
	"encoding/json"
	"io"
	"strings"
)

type Template struct {
	Title      string    `json:"title"`
	ID         int       `json:"id"`
	Text       string    `json:"text"`
	Variants   []Variant `json:"-"`
	Parent     string    `json:"-"`
	SpeechPart string    `json:"-"`
}
type Variant struct {
	Form     Form
	Patterns []string
}

func ParseTemplates(r io.Reader) ([]Template, error) {
	var templates []Template
	if err := json.NewDecoder(r).Decode(&templates); err != nil {
		return nil, err
	}

	var handledTemplates []Template
	for _, t := range templates {
		if t.detectSpeechPart(); t.SpeechPart == "" {
			continue
		}
		if t.isPureSpeechPart() {
			t.Variants = []Variant{{Form: pureForm}}
			handledTemplates = append(handledTemplates, t)
			continue
		}

		if t.parseTemplate(); len(t.Variants) == 0 {
			continue
		}
		handledTemplates = append(handledTemplates, t)
	}
	//for _, t := range handledTemplates {
	//fmt.Println(t.Title, t.Variants)
	//}

	return handledTemplates, nil
}

func (t *Template) GetFormVariants(bases map[string]string) []string {
	var variants []string
	for _, v := range t.Variants {
		for _, f := range v.Patterns {
			prefix, baseName, suffix := GetFormParts(f)

			base, ok := bases[baseName]
			if !ok {
				continue
			}

			variants = append(variants, prefix+base+suffix)
		}
	}
	return variants
}

func GetFormParts(s string) (string, string, string) {
	const (
		openTag  = "{{{"
		closeTag = "}}}"
	)
	start := strings.Index(s, openTag)
	if start == -1 {
		return "", "", ""
	}

	end := strings.Index(s, closeTag)
	if end == -1 {
		return "", "", ""
	}
	return s[:start], s[start+len(openTag) : end], s[end+len(closeTag):]
}
