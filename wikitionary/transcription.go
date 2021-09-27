package wikitionary

import (
	"strings"
	"unicode"
)

func getTranscription(s string) string {
	if !strings.HasPrefix(s, "transcription-ru") {
		return ""
	}
	fields := strings.Split(s, "|")
	if len(fields) < 2 {
		return ""
	}
	s = fields[1]

	if s == "" {
		return ""
	} else if strings.ContainsRune(s, ' ') {
		return ""
	}

	s = strings.Replace(s, `́`, ``, -1) // знак ударения
	s = strings.ToLower(s)
	for _, r := range []rune(s) {
		if !unicode.Is(unicode.Cyrillic, r) {
			return ""
		}
	}
	return s
}
