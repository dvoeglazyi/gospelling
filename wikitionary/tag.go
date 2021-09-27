package wikitionary

import (
	"bufio"
	"strings"
)

func readTag(r *bufio.Reader, reminder string) (string, string, error) {
	const (
		openTag  = "{{"
		closeTag = "}}"
	)
	var result string
	openedTagLevel := -1
	for {
		if reminder == "" {
			var err error
			if reminder, err = r.ReadString('\n'); err != nil {
				return "", "", err
			}
		}
		if openedTagLevel < 0 {
			if index := strings.Index(reminder, openTag); index >= 0 {
				openedTagLevel++
				reminder = reminder[index:]

				if indexEnd := strings.Index(reminder, closeTag); indexEnd >= 0 {

					inner := reminder[:indexEnd]
					if openedTagLevel += strings.Count(inner, openTag); openedTagLevel == 0 {
						return reminder[indexEnd:], inner, nil
					}
				}
				result = reminder
			}
			reminder = ""
			continue
		}
		if indexEnd := strings.Index(reminder, closeTag); indexEnd >= 0 {
			openedTagLevel--

			inner := reminder[:indexEnd]
			if openedTagLevel += strings.Count(inner, openTag); openedTagLevel == 0 {
				result += "\n" + inner
				return reminder[indexEnd:], result, nil
			}
		}
		result += "\n" + reminder
		reminder = ""
	}
}

func trimTag(s string) string {
	s = strings.TrimPrefix(s, "{{")
	return strings.TrimSuffix(s, "}}")
}
