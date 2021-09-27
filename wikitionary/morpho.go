package wikitionary

import (
	"strings"
)

func GetWordBasesFromTag(s string) map[string]string {
	s = strings.Replace(s, "\n", "", -1)
	tagParts := strings.Split(s, "|")
	bases := make(map[string]string)
	for _, part := range tagParts {
		if !strings.HasPrefix(part, "основа") {
			continue
		}
		fields := strings.Split(part, "=")
		if len(fields) != 2 {
			continue
		}
		key := strings.TrimSpace(fields[0])
		bases[key] = fields[1]
	}

	return bases
}
