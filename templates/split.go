package templates

import (
	"strings"
)

type Tag struct {
	open, close string
}

// Если среди тэгов встречаются такие, которые могут встречаться в составе других
// (пример: тэг "{" в составе "{{"),
// то для того, чтобы функция index правильно работала,
// они должны быть упорядочены так, чтобы тэги, которые в составе, были в конце.
var tags = []Tag{{open: "{{{", close: "}}}"}, {open: "{{", close: "}}"}, {open: "[[", close: "]]"}, {open: "{", close: "}"}}

func index(s string, sep string, closeTagID int) (tagID int, isOpen bool, index int, length int) {
	for i := range s {
		if strings.HasPrefix(s[i:], sep) {
			return -1, false, i, len(sep)
		} else if closeTagID != -1 && strings.HasPrefix(s[i:], tags[closeTagID].close) {
			return closeTagID, false, i, len(tags[closeTagID].close)
		}
		for tagID, tag := range tags {
			if strings.HasPrefix(s[i:], tag.open) {
				return tagID, true, i, len(tag.open)
			}
		}
	}
	return -1, false, -1, 0
}

// Split разделяет строку подобно strings.Split, но с учётом экранирующих тэгов (внутри них разделители игнорируются).
func Split(s string, sep string) []string {
	var tagIDs []int
	n := strings.Count(s, sep) + 1
	a := make([]string, n)
	n--
	i := 0
	for i < n {
		currentTagID := -1
		if len(tagIDs) > 0 {
			currentTagID = tagIDs[len(tagIDs)-1]
		}

		tagID, isOpen, pos, length := index(s, sep, currentTagID)
		if pos < 0 {
			break
		}
		if tagID >= 0 {
			if isOpen {
				// найден открывающий тэг
				tagIDs = append(tagIDs, tagID)
			} else if len(tagIDs) > 0 && tagIDs[len(tagIDs)-1] == tagID {
				// найден подходящий закрывающий тэг
				tagIDs = tagIDs[:len(tagIDs)-1]
			}
		}

		if len(tagIDs) == 0 && tagID == -1 {
			// найден разделитель
			a[i] += s[:pos]
			i++
		} else {
			// либо найден тэг, либо разделитель экранирован
			a[i] += s[:pos+length]
		}
		s = s[pos+length:]

	}

	a[i] += s
	return a[:i+1]
}
